/**********************************************
** @Des: 学生基础信息
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package models

import (
	"github.com/astaxie/beego/orm"
)

// BaseStudent 学生基本信息表
type BaseStudent struct {
	ID           int    `json:"id" orm:"column(id);auto"`
	Name         string `json:"name" orm:"column(name)" description:"姓名"`
	Idcard       string `json:"idcard" orm:"column(idcard)" description:"身份证id"`
	Batch        string `json:"batch" orm:"column(batch)" description:"年级"`
	School       string `json:"school" orm:"column(school)" description:"学校"`
	Course       string `json:"course" orm:"column(course)" description:"课程组合"`
	BaseClass    string `json:"baseclass" orm:"column(base_class)" description:"行政班"`
	SubjectClass string `json:"subjectclass" orm:"column(subject_class)" description:"学科班"`
	Level        string `json:"level" orm:"column(level)" description:"类型"`
}

// TableName 表名
func (t *BaseStudent) TableName() string {
	return "base_student"
}

func init() {
	orm.RegisterModel(new(BaseStudent))
}

// AddBaseStudent insert a new BaseStudent into database and returns
// last inserted Id on success.
func AddBaseStudent(m *BaseStudent) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetStudentList 获学生列表
func GetStudentList(page, pageSize int, filters ...interface{}) ([]*BaseStudent, int64) {
	offset := (page - 1) * pageSize
	list := make([]*BaseStudent, 0)
	query := orm.NewOrm().QueryTable("base_student")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

// GetStudentCount is ...
// the record to be updated doesn't exist
func GetStudentCount(batch string) (int64, error) {
	return orm.NewOrm().QueryTable("base_student").Filter("Batch", batch).Count()
}
