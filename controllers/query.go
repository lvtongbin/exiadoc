/**********************************************
** @Des: 查看数据
** @Author: lvtongbin
** @Date:   2019-9-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-10-09 18:50:41
***********************************************/

package controllers

import (
	"strconv"
	"strings"
	"trism/models"
	"trism/myredis"
)

// QueryController is Controller of query
type QueryController struct {
	BaseController
}

// Index 登录入口
func (queryController *QueryController) Index() {
	// 获取年级
	batchList, _ := models.GetBatchList(1, 99)

	queryController.Data["gradeList"] = batchList
	queryController.TplName = "query/index.html"
}

// SelectSchoolAndExam is ...
// @Title SelectSchoolAndExam
// @Description 获取学校和考试
// @router /selectSchoolAndExam [get]
func (queryController *QueryController) SelectSchoolAndExam() {
	grade := strings.TrimSpace(queryController.GetString("grade", ""))
	if grade == "" {
		queryController.ajaxMsg("请选择正确的年级", MessageError)
	}

	filters := make([]interface{}, 0)
	filters = append(filters, "Batch", grade)
	examList, _ := models.GetExamList(1, 999, filters...)

	head := grade[0:6]
	schoolfilters := make([]interface{}, 0)
	schoolfilters = append(schoolfilters, "name__icontains", head)
	schoolList, _ := models.GetSchoolList(1, 999, schoolfilters...)

	data := map[string]interface{}{"exams": examList, "schools": schoolList}
	queryController.ajaxList("成功", MessageOK, 0, data)
}

// SelectClass is ...
// @Title SelectClass
// @Description 获取班级
// @router /selectClass [get]
func (queryController *QueryController) SelectClass() {
	grade := strings.TrimSpace(queryController.GetString("grade", ""))
	school := strings.TrimSpace(queryController.GetString("school", ""))
	if grade == "" || school == "" {
		queryController.ajaxMsg("请选择正确的年级或学校", MessageError)
	}

	filters := make([]interface{}, 0)
	filters = append(filters, "Batch", grade)
	filters = append(filters, "School", school)
	classList, count := models.GetBaseClasssList(1, 999, filters...)

	queryController.ajaxList("成功", MessageOK, count, classList)
}

// QueryExam is ...
// @Title QueryExam
// @Description 查询成绩
// @router /queryExam [get]
func (queryController *QueryController) QueryExam() {
	page, _ := queryController.GetInt("page", 1)
	limit, _ := queryController.GetInt("limit", 50)

	grade := strings.TrimSpace(queryController.GetString("grade", ""))
	exam := strings.TrimSpace(queryController.GetString("exam", ""))
	school := strings.TrimSpace(queryController.GetString("school", ""))
	class := strings.TrimSpace(queryController.GetString("class", ""))
	if exam == "" {
		queryController.ajaxMsg("请选择正确的考试", MessageError)
	}
	if grade == "" || school == "" || class == "" {
		queryController.ajaxMsg("请选择正确的年级、学校或班级", MessageError)
	}
	batch, err := models.GetItemsBatch(grade)
	if err != nil {
		queryController.ajaxMsg("获取年级失败", MessageError)
	}

	examScores, count := models.GetExamScoreList(page, limit, grade, batch.Code, exam, school, class)

	queryController.ajaxList("成功", MessageOK, count, examScores)
}

// GetSummary is ...
// @Title GetSummary
// @Description 获取班级
// @router /getSummary [get]
func (queryController *QueryController) GetSummary() {
	grade := strings.TrimSpace(queryController.GetString("grade", ""))
	exam := strings.TrimSpace(queryController.GetString("exam", ""))
	school := strings.TrimSpace(queryController.GetString("school", ""))
	class := strings.TrimSpace(queryController.GetString("class", ""))
	if exam == "" {
		queryController.ajaxMsg("请选择正确的考试", MessageError)
	}
	if grade == "" || school == "" || class == "" {
		queryController.ajaxMsg("请选择正确的年级、学校或班级", MessageError)
	}

	subjectCount := 3
	if strings.Contains(grade, "中学") {
		subjectCount = 8
	}

	key := ""
	if school == "all" {
		// 获取年级概述
		key = "score_" + exam + "_total"

	} else if class == "all" {
		// 获取学校概叙
		organize, _ := models.GetBaseOrganizeByName(school)
		key = "score_" + exam + "_" + organize.Code + "_total"
	} else {
		// 获取班级概述
		organize, _ := models.GetBaseOrganizeByName(school)
		class, _ := models.GetBaseClasssByName(class, school)
		key = "score_" + exam + "_" + organize.Code + "_" + class.Code + "_total"
	}

	if key == "" {
		queryController.ajaxMsg("参数错误", MessageError)
	}
	eachMembers, _ := myredis.GetEachScore(key)
	x := make([]string, 0, 10)
	y := make([]int, 10, 10)

	rx := []string{"优", "良", "及格", "不及格"}
	ry := make([]int, 4, 4)

	for i := 0; i < 10; i++ {
		a := strconv.Itoa(10*(9-i)*subjectCount) + "+"
		x = append(x, a)
	}
	for _, v := range eachMembers {
		for i := 0; i < 100*subjectCount+1; i++ {
			if v.Score == i {
				if i < 10*subjectCount {
					y[9] += v.Total
					ry[3] += v.Total
				} else if i < 20*subjectCount {
					y[8] += v.Total
					ry[3] += v.Total
				} else if i < 30*subjectCount {
					y[7] += v.Total
					ry[3] += v.Total
				} else if i < 40*subjectCount {
					y[6] += v.Total
					ry[3] += v.Total
				} else if i < 50*subjectCount {
					y[5] += v.Total
					ry[3] += v.Total
				} else if i < 60*subjectCount {
					y[4] += v.Total
					ry[3] += v.Total
				} else if i < 70*subjectCount {
					y[3] += v.Total
					ry[2] += v.Total
				} else if i < 80*subjectCount {
					y[2] += v.Total
					if i < 75*subjectCount {
						ry[2] += v.Total
					} else {
						ry[1] += v.Total
					}
				} else if i < 90*subjectCount {
					y[1] += v.Total
					ry[1] += v.Total
				} else {
					y[0] += v.Total
					ry[0] += v.Total
				}
				break
			}
		}
	}

	ryy := make([]map[string]interface{}, 0, 4)
	for i := 0; i < 4; i++ {
		yy := map[string]interface{}{"value": ry[i], "name": rx[i]}
		ryy = append(ryy, yy)
	}

	score := map[string]interface{}{"x": x, "y": y}
	scoreRing := map[string]interface{}{"x": rx, "y": ryy}
	data := map[string]interface{}{"score": score, "scoreRing": scoreRing}

	queryController.ajaxList("成功", MessageOK, 0, data)
}
