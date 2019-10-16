package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// GetCodeString used to generate code strings.
func GetCodeString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomString used to generate random strings.
func GetRandomString(lens int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetGradeRandom 获取成绩
func GetGradeRandom() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(10 * time.Millisecond))
	return r.Intn(101)
}

// GetNameRandom 获取姓名
func GetNameRandom() string {
	xings := []string{
		"王", "李", "张", "刘", "陈", "杨", "黄", "赵", "吴", "周", "徐", "孙", "马", "朱", "胡", "郭", "何", "高", "林", "罗",
		"郑", "梁", "谢", "宋", "唐", "许", "韩", "冯", "邓", "曹", "彭", "曾", "肖", "田", "董", "袁", "潘", "于", "蒋", "蔡",
		"余", "杜", "叶", "程", "苏", "魏", "吕", "丁", "任", "沈", "姚", "卢", "姜", "崔", "钟", "谭", "陆", "汪", "范", "金",
		"石", "廖", "贾", "夏", "韦", "付", "方", "白", "邹", "孟", "熊", "秦", "邱", "江", "尹", "薛", "闫", "段", "雷", "侯",
		"龙", "史", "陶", "黎", "贺", "顾", "毛", "郝", "龚", "邵", "万", "钱", "严", "覃", "武", "戴", "莫", "孔", "向", "汤",
	}
	mings := []string{
		"伟", "刚", "勇", "毅", "俊", "峰", "强", "军", "平", "保", "东", "文", "辉", "力",
		"明", "永", "健", "世", "广", "志", "义", "兴", "良", "海", "山", "仁", "波", "宁",
		"贵", "福", "生", "龙", "元", "全", "国", "胜", "学", "祥", "才", "发", "武", "新",
		"利", "清", "飞", "彬", "富", "顺", "信", "子", "杰", "涛", "昌", "成", "康", "星",
		"光", "天", "达", "安", "岩", "中", "茂", "进", "林", "有", "坚", "和", "彪", "博",
		"诚", "先", "敬", "震", "振", "壮", "会", "思", "群", "豪", "心", "邦", "承", "乐",
		"绍", "功", "松", "善", "厚", "庆", "磊", "民", "友", "裕", "河", "哲", "江", "超",
		"浩", "亮", "政", "谦", "亨", "奇", "固", "之", "轮", "翰", "朗", "伯", "宏", "言",
		"若", "鸣", "朋", "斌", "梁", "栋", "维", "启", "克", "伦", "翔", "旭", "鹏", "泽",
		"晨", "辰", "士", "以", "建", "家", "致", "树", "炎", "德", "行", "时", "泰", "盛",
		"秀", "娟", "英", "华", "慧", "巧", "美", "娜", "静", "淑", "惠", "珠", "翠", "雅",
		"芝", "玉", "萍", "红", "娥", "玲", "芬", "芳", "燕", "彩", "春", "菊", "兰", "凤",
		"洁", "梅", "琳", "素", "云", "莲", "真", "环", "雪", "荣", "爱", "妹", "霞", "香",
		"月", "莺", "媛", "艳", "瑞", "凡", "佳", "嘉", "琼", "勤", "珍", "贞", "莉", "桂",
		"娣", "叶", "璧", "璐", "娅", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛",
		"青", "倩", "婷", "姣", "婉", "娴", "瑾", "颖", "露", "瑶", "怡", "婵", "雁", "蓓",
		"纨", "仪", "荷", "丹", "蓉", "眉", "君", "琴", "蕊", "薇", "菁", "梦", "岚", "苑",
		"筠", "柔", "竹", "霭", "凝", "晓", "欢", "霄", "枫", "芸", "菲", "寒", "欣", "滢",
		"伊", "亚", "宜", "可", "姬", "舒", "影", "荔", "枝", "思", "丽", "秀", "飘", "育",
		"馥", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛", "青", "倩", "婷", "宁",
		"蓓", "纨", "苑", "婕", "馨", "瑗", "琰", "韵", "融", "园", "艺", "咏", "卿", "聪",
		"澜", "纯", "毓", "悦", "昭", "冰", "爽", "琬", "茗", "羽", "希",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	xing := xings[r.Intn(100)]
	var ming string
	b := r.Intn(10)
	if b%2 == 0 {
		ming = mings[r.Intn(304)]
	} else {
		ming = mings[r.Intn(304)] + mings[r.Intn(305)]
	}
	return xing + ming
}

// GetSchoolRandom 获取学校、年级、班级
func GetSchoolRandom() (string, string, string) {
	xschool := []string{"A县第一小学", "B县第一小学", "B县第二小学", "C县第一小学", "D县第一小学"}
	mschool := []string{"A县第一中学", "A县第二中学", "B县第一中学", "B县第二中学", "C县第一中学", "C县第二中学"}
	sclass := []string{"一班", "二班", "三班", "四班"}

	var school string
	var nianji string
	var class string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := r.Intn(8)
	if b < 5 {
		school = xschool[b]
		nianji = "小学" + strconv.Itoa(15+b) + "级"
		class = strconv.Itoa(15+b) + "级" + sclass[r.Intn(4)]
	} else {
		school = mschool[r.Intn(6)]
		nianji = "中学" + strconv.Itoa(12+b) + "级"
		class = strconv.Itoa(12+b) + "级" + sclass[r.Intn(4)]
	}
	return school, nianji, class
}
