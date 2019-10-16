package controllers

import (
	"fmt"
	"strings"
	"time"
	"trism/longpoll"
	"trism/models"
	"trism/myredis"
	"trism/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// BaseExamController operations for BaseExam
type BaseExamController struct {
	beego.Controller
}

// SimulateExam is ...
// @Title SimulateExam
// @Description 添加新的考试
// @router /simulateExam [post]
func (c *BaseExamController) SimulateExam() {
	grade := c.GetString("grade")
	examname := c.GetString("examname")
	token := c.GetString("token")

	if grade == "" || examname == "" {
		c.Data["json"] = map[string]interface{}{"code": 1001, "errmsg": "参数错误"}
		c.ServeJSON()
	}

	batch, _ := models.GetItemsBatch(grade)

	subjects := ""
	if strings.Contains(grade, "小学") {
		subjects = "语文;数学;英语"
	} else {
		subjects = "语文;数学;英语;物理;化学;政治;历史;生物;地理"
	}

	code := utils.GetCodeString(10)
	baseExam := &models.BaseExam{
		Code:     code,
		Batch:    grade,
		Subjects: subjects,
		Name:     examname,
		Datetime: time.Now().Unix(),
	}
	if _, err := models.AddBaseExam(baseExam); err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1002, "errmsg": "添加考试失败"}
		c.ServeJSON()
	}

	go simulateScore(batch, code, token)

	c.Data["json"] = map[string]int64{"code": 0}
	c.ServeJSON()
}

func simulateScore(grade *models.ItemsBatch, code, token string) {
	count, _ := models.GetStudentCount(grade.Name)

	times := int(count/20 + 1)

	for i := 0; i < times; i++ {
		page := i + 1
		pageSize := 0
		if page < times {
			pageSize = 20
		} else {
			pageSize = int(count) - i*20
		}
		if pageSize == 0 {
			break
		}
		filters := make([]interface{}, 0)
		filters = append(filters, "Batch", grade.Name)
		baseStudents, _ := models.GetStudentList(page, pageSize, filters...)

		for _, v := range baseStudents {
			school := ""
			class := ""
			// 获取学校code,可以优化成从缓存中取
			if baseOrganize, err := models.GetBaseOrganizeByName(v.School); err == nil {
				school = baseOrganize.Code
			}
			// 获取年级code,可以优化成从缓存中取
			if baseClass, err := models.GetBaseClasssByName(v.BaseClass, v.School); err == nil {
				class = baseClass.Code
			}
			if strings.Contains(grade.Name, "小学") {
				chinese := utils.GetGradeRandom()
				math := utils.GetGradeRandom()
				english := utils.GetGradeRandom()
				// 登记成绩
				testScores := &models.ExamScores{
					Grade:     grade.Code,
					StudentID: v.Idcard,
					ExamID:    code,
					Chinese:   chinese,
					Math:      math,
					English:   english,
					Politics:  -1,
					History:   -1,
					Chemistry: -1,
					Physics:   -1,
					Geography: -1,
					Biology:   -1,
					Total:     chinese + math + english,
				}
				models.AddExamScores(testScores)
				// 添加日志
				data := fmt.Sprintf("id=%s&school=%s&class=%s&chinese=%d&math=%d&english=%d",
					code, school, class, chinese, math, english)
				logs.Info(data)

			} else {
				chinese := utils.GetGradeRandom()
				math := utils.GetGradeRandom()
				english := utils.GetGradeRandom()
				politics := utils.GetGradeRandom()
				history := utils.GetGradeRandom()
				chemistry := utils.GetGradeRandom()
				physics := utils.GetGradeRandom()
				gengraphy := utils.GetGradeRandom()
				biology := utils.GetGradeRandom()

				// 登记成绩
				testScores := &models.ExamScores{
					Grade:     grade.Code,
					StudentID: v.Idcard,
					ExamID:    code,
					Chinese:   chinese,
					Math:      math,
					English:   english,
					Politics:  politics,
					History:   history,
					Chemistry: chemistry,
					Physics:   physics,
					Geography: gengraphy,
					Biology:   biology,
					Total:     chinese + math + english + politics + history + chemistry + physics + gengraphy + biology,
				}
				models.AddExamScores(testScores)

				// 添加日志
				data := fmt.Sprintf("id=%s&school=%s&class=%s&chinese=%d&math=%d&english=%d&politics=%d&history=%d&chemistry=%d&physics=%d&gengraphy=%d&biology=%d",
					code, school, class, chinese, math, english, politics, history, chemistry, physics, gengraphy, biology)
				logs.Info(data)
			}
		}
		//添加进度
		complete := i*20 + pageSize

		// 获取参数
		key := "score_" + code + "_chinese"
		eachMembers, _ := myredis.GetEachScore(key)
		x := make([]int, 0, 101)
		y := make([]int, 101, 101)

		rx := []string{"优", "良", "及格", "不及格"}
		ry := make([]int, 4, 4)

		for i := 0; i < 101; i++ {
			x = append(x, i)
		}
		for _, v := range eachMembers {
			for i := 0; i < 101; i++ {
				if v.Score == i {
					y[i] = v.Total
					if i < 60 {
						ry[3] += v.Total
					} else if i < 75 {
						ry[2] += v.Total
					} else if i < 90 {
						ry[1] += v.Total
					} else {
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

		chineseScore := map[string]interface{}{"x": x, "y": y}
		chineseScoreRing := map[string]interface{}{"x": rx, "y": ryy}
		data := map[string]interface{}{"complete": complete, "total": count, "chineseScore": chineseScore, "chineseScoreRing": chineseScoreRing}

		longpoll.Publish(token, data)
	}
}
