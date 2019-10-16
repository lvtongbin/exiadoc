package statistics

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"github.com/mediocregopher/radix.v2/pool"
)

func init() {
	// 初始化log日志
	logs.SetLogger(logs.AdapterFile, `{"filename":"trism_grade.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":7,"color":false}`)

	// 初始化channel
	var logChannel = make(chan string, 3*routineNum)
	var gradeChannel = make(chan GradeData, routineNum)
	var storageChannel = make(chan AnalyzData, 2*routineNum)

	redisPool, err := pool.New("tcp", "localhost:6379", 2*routineNum)
	if err != nil {
		fmt.Println("Redis pool created failed:", err)
		return
	}
	go func() {
		for {
			redisPool.Cmd("PING")
			time.Sleep(3 * time.Second)
		}
	}()

	// 消费日志
	go readFileLinebyLine(logChannel)

	// 日志处理
	for i := 0; i < routineNum; i++ {
		go logConsumer(logChannel, gradeChannel)
	}

	// 日志解析
	go analyzCounter(gradeChannel, storageChannel)

	// 创建存储器
	for i := 0; i < routineNum; i++ {
		go dataStorage(storageChannel, redisPool)
	}

	// 开启任务
	startDailyJob(redisPool)
}

// readFileLinebyLine 从log日志中一行一行的读取信息
func readFileLinebyLine(logChannel chan string) error {
	// 启动tail
	tails, err := tail.TailFile("trism_grade.log", tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return err
	}

	for {
		msg, ok := <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filenam:%s\n", "project.log")
			time.Sleep(100 * time.Millisecond)
			continue
		}
		logChannel <- msg.Text
	}
}

func logConsumer(logChannel chan string, gradeChannel chan GradeData) error {
	for logstr := range logChannel {
		if len(logstr) == 0 {
			continue
		}
		logs := strings.Split(logstr, "]")
		if len(logs) < 3 {
			continue
		}
		data, err := newGradeData(logs[2])
		if err != nil {
			continue
		}
		// 统计
		gradeChannel <- *data
	}
	return nil
}

func analyzCounter(gradeChannel chan GradeData, storageChannel chan AnalyzData) error {
	for data := range gradeChannel {
		refValue := reflect.ValueOf(data)
		refType := reflect.TypeOf(data)

		fieldCount := refValue.NumField()
		for i := 0; i < fieldCount; i++ {
			fieldValue := refValue.Field(i)
			fieldType := refType.Field(i)
			if v, ok := fieldValue.Interface().(int); ok && v >= 0 {
				fieldTag := fieldType.Tag.Get("key")
				storageChannel <- AnalyzData{data.ID, data.School, data.Class, fieldTag, v}
			}
		}
	}
	return nil
}

func dataStorage(storageChannel chan AnalyzData, redisPool *pool.Pool) {
	for block := range storageChannel {
		setKeys := []string{
			// 全县(联考)统计
			fmt.Sprintf("%s%s_%s", prefix, block.ID, block.Subject),
			// 学校统计
			fmt.Sprintf("%s%s_%s_%s", prefix, block.ID, block.School, block.Subject),
			// 班级统计
			fmt.Sprintf("%s%s_%s_%s_%s", prefix, block.ID, block.School, block.Class, block.Subject),
		}

		for _, key := range setKeys {
			ret, err := redisPool.Cmd("ZINCRBY", key, 1, block.Score).Int()
			if err != nil || ret <= 0 {
				fmt.Println("DataStorage redis storage error.", block.ID, block.School, block.Class)
			}
			// 保存一个月
			redisPool.Cmd("EXPIRE", key, 35*86400).Int()
		}
	}
}
