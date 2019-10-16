/**********************************************
** @Des: 考试信息
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package models

import (
	"github.com/astaxie/beego/orm"
)

// BaseExam 考试信息基础表
type BaseExam struct {
	ID       int    `json:"id" orm:"column(id);auto"`
	Code     string `json:"code" orm:"column(code)" description:"考试编码"`
	Batch    string `json:"batch" orm:"column(batch)" description:"年级"`
	Datetime int64  `json:"datetime" orm:"column(datetime)" description:"时间"`
	Name     string `json:"name" orm:"column(name)" description:"考试名称"`
	Subjects string `json:"subjects" orm:"column(subjects)" description:"科目列表"`
}

// TableName 表名
func (t *BaseExam) TableName() string {
	return "base_exam"
}

func init() {
	orm.RegisterModel(new(BaseExam))
}

// AddBaseExam insert a new BaseExam into database and returns
// last inserted Id on success.
func AddBaseExam(m *BaseExam) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExamList 获取考试列表
func GetExamList(page, pageSize int, filters ...interface{}) ([]*BaseExam, int64) {
	offset := (page - 1) * pageSize
	list := make([]*BaseExam, 0)
	query := orm.NewOrm().QueryTable("base_exam")
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
