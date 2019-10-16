package statistics

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

const (
	infoData    = "I"
	warningData = "W"
	errorData   = "E"
	routineNum  = 5
	prefix      = "score_"
)

// GradeData 成绩表
type GradeData struct {
	ID        string `key:"id"`
	School    string `key:"school"`
	Class     string `key:"class"`
	Chinese   int    `key:"chinese"`
	Math      int    `key:"math"`
	English   int    `key:"english"`
	Politics  int    `key:"politics"`
	History   int    `key:"history"`
	Chemistry int    `key:"chemistry"`
	Physics   int    `key:"physics"`
	Gengraphy int    `key:"gengraphy"`
	Biology   int    `key:"biology"`
	Sports    int    `key:"sports"`
	Total     int    `key:"total"`
}

// AnalyzData 解析数据
type AnalyzData struct {
	ID      string // 考试id
	School  string // 学校
	Class   string // 班级
	Subject string // 科目
	Score   int
}

func newGradeData(log string) (*GradeData, error) {
	urlInfo, err := url.Parse("http://localhost/?" + strings.TrimSpace(log))
	if err != nil {
		return nil, err
	}
	data := urlInfo.Query()
	d := &GradeData{
		ID:     data.Get("id"),
		School: data.Get("school"),
		Class:  data.Get("class"),
	}
	if d.ID == "" || d.School == "" || d.Class == "" {
		return nil, errors.New("<gradeData>Param is illegal")
	}
	d.Chinese = getItemGrade(data, "chinese")
	if d.Chinese > -1 {
		d.Total = d.Total + d.Chinese
	}
	d.Math = getItemGrade(data, "math")
	if d.Math > -1 {
		d.Total = d.Total + d.Math
	}
	d.English = getItemGrade(data, "english")
	if d.English > -1 {
		d.Total = d.Total + d.English
	}
	d.Politics = getItemGrade(data, "politics")
	if d.Politics > -1 {
		d.Total = d.Total + d.Politics
	}
	d.History = getItemGrade(data, "history")
	if d.History > -1 {
		d.Total = d.Total + d.History
	}
	d.Chemistry = getItemGrade(data, "chemistry")
	if d.Chemistry > -1 {
		d.Total = d.Total + d.Chemistry
	}
	d.Physics = getItemGrade(data, "physics")
	if d.Physics > -1 {
		d.Total = d.Total + d.Physics
	}
	d.Gengraphy = getItemGrade(data, "gengraphy")
	if d.Gengraphy > -1 {
		d.Total = d.Total + d.Gengraphy
	}
	d.Biology = getItemGrade(data, "biology")
	if d.Biology > -1 {
		d.Total = d.Total + d.Biology
	}
	d.Sports = getItemGrade(data, "sports")
	if d.Sports > -1 {
		d.Total = d.Total + d.Sports
	}

	return d, nil
}

func getItemGrade(value url.Values, key string) int {
	v := value.Get(key)
	if v == "" {
		return -1
	}
	grade, err := strconv.Atoi(v)
	if err != nil {
		return -1
	}
	return grade
}
