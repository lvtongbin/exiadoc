package statistics

import (
	cron "fdmonitor/crons"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/mediocregopher/radix.v2/pool"
)

// CronJob 定时任务
type CronJob struct {
	RedisPool *pool.Pool
}

func startDailyJob(redisPool *pool.Pool) {
	mainCron := cron.New()
	cronJob := &CronJob{redisPool}
	mainCron.AddJob("@daily", cronJob)
	b, _ := beego.AppConfig.Bool("taskAtOnce")
	if b {
		cronJob.Run()
	}
	mainCron.Start()
}

// Run 执行定时任务
func (cronJob CronJob) Run() {
	fmt.Println("CronJob is Run!")
	// ct := time.Now().Add(60 * time.Second) // 防止执行时还在之前一天

}
