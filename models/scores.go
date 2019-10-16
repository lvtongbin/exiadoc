/**********************************************
** @Des: 学生成绩信息
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// ExamScores 学生成绩表
type ExamScores struct {
	Name      string `json:"name"`
	Grade     string `json:"grade"`
	ExamID    string `json:"exam_id"`
	StudentID string `json:"student_id" orm:"column(student_id)"`
	Biology   int    `json:"biology"`
	Chemistry int    `json:"chemistry"`
	Chinese   int    `json:"chinese"`
	English   int    `json:"english"`
	Geography int    `json:"geography"`
	History   int    `json:"history"`
	Math      int    `json:"math"`
	Physics   int    `json:"physics"`
	Politics  int    `json:"politics"`
	Sports    int    `json:"sports"`
	Total     int    `json:"total"`
	Sort      string `json:"sort"`
}

func scoreTableName(code string) string {
	return "exam_scores_" + code
}

// AddExamScores insert a new TestScores into database and returns
// last inserted Id on success.
func AddExamScores(examScores *ExamScores) error {
	sql := fmt.Sprintf("INSERT INTO %s(student_id, exam_id, chinese, english, math, politics, history, chemistry, physics, geography, biology, sports, total) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", scoreTableName(examScores.Grade))
	_, err := orm.NewOrm().Raw(sql, examScores.StudentID, examScores.ExamID, examScores.Chinese,
		examScores.English, examScores.Math, examScores.Politics, examScores.History, examScores.Chemistry,
		examScores.Physics, examScores.Geography, examScores.Biology, examScores.Sports, examScores.Total).Exec()
	return err
}

// GetExamScoreList 获取成绩列表
func GetExamScoreList(page, pageSize int, batch, grade, exam, school, class string) ([]ExamScores, int64) {
	offset := (page - 1) * pageSize
	var examScoreList []ExamScores
	var count int64
	if school == "all" {
		// 查看某次考试全部成绩
		sql := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE exam_id=?", scoreTableName(grade))
		err := orm.NewOrm().Raw(sql, exam).QueryRow(&count)

		if err == nil && count > 0 {
			sql = fmt.Sprintf(`SELECT a.name,b.student_id,b.chinese,b.english,b.math,b.politics,b.history,b.chemistry,b.physics,b.geography,b.biology,b.sports,b.total,b.sort 
				FROM base_student AS a,%s AS b WHERE a.idcard=b.student_id AND b.exam_id=? ORDER BY b.total DESC LIMIT ?,? `,
				scoreTableName(grade))
			orm.NewOrm().Raw(sql, exam, offset, pageSize).QueryRows(&examScoreList)
		}
	} else if class == "all" {
		// 查看学校某次考斯的全部成绩
		sql := "SELECT COUNT(*) FROM base_student WHERE school=? AND batch=?"
		err := orm.NewOrm().Raw(sql, school, batch).QueryRow(&count)
		if err == nil && count > 0 {
			sql = fmt.Sprintf(`SELECT a.name,b.student_id,b.chinese,b.english,b.math,b.politics,b.history,b.chemistry,b.physics,b.geography,b.biology,b.sports,b.total,b.sort 
				FROM base_student AS a,%s AS b WHERE a.idcard=b.student_id AND b.exam_id=? AND a.school=? ORDER BY b.total DESC LIMIT ?,? `,
				scoreTableName(grade))
			orm.NewOrm().Raw(sql, exam, school, offset, pageSize).QueryRows(&examScoreList)
		}
	} else {
		// 查看班级某次考斯的全部成绩
		sql := "SELECT COUNT(*) FROM base_student WHERE school=? AND batch=? AND base_class=?"
		err := orm.NewOrm().Raw(sql, school, batch, class).QueryRow(&count)
		if err == nil && count > 0 {
			sql = fmt.Sprintf(`SELECT a.name,b.student_id,b.chinese,b.english,b.math,b.politics,b.history,b.chemistry,b.physics,b.geography,b.biology,b.sports,b.total,b.sort 
				FROM base_student AS a,%s AS b WHERE a.idcard=b.student_id AND b.exam_id=? AND a.school=? AND a.base_class=? ORDER BY b.total DESC LIMIT ?,? `,
				scoreTableName(grade))
			orm.NewOrm().Raw(sql, exam, school, class, offset, pageSize).QueryRows(&examScoreList)
		}
	}
	return examScoreList, count
}
