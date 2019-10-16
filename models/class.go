/**********************************************
** @Des: 班级，包括行政班和学科班
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package models

import "github.com/astaxie/beego/orm"

// BaseClasss 行政班基础信息表
type BaseClasss struct {
	ID     int    `json:"id" orm:"column(id);auto"`
	Code   string `json:"code" orm:"column(code)" description:"编码"`
	Batch  string `json:"batch" orm:"column(batch)" description:"年级"`
	Name   string `json:"name" orm:"column(name)" description:"行政班名"`
	School string `json:"school" orm:"column(school)" description:"学校"`
}

// SubjectClass 学科班表
type SubjectClass struct {
	ID      int    `json:"id" orm:"column(id);auto" description:"自增id"`
	Code    string `json:"code" orm:"column(code)" description:"编码"`
	Batch   string `json:"batch" orm:"column(batch)" description:"年级"`
	School  string `json:"school" orm:"column(school)" description:"学校"`
	Subject string `json:"subject" orm:"column(subject)" description:"学科"`
}

// TableName 表名
func (t *BaseClasss) TableName() string {
	return "base_classs"
}

// TableName 表名
func (t *SubjectClass) TableName() string {
	return "subject_class"
}

func init() {
	orm.RegisterModel(new(SubjectClass), new(BaseClasss))
}

// AddBaseClasss 新增行政班级
func AddBaseClasss(m *BaseClasss) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBaseClasssByName 获取行政班级基本信息
func GetBaseClasssByName(name, school string) (*BaseClasss, error) {
	baseClasss := new(BaseClasss)
	err := orm.NewOrm().QueryTable(baseClasss).Filter("Name", name).Filter("School", school).One(baseClasss)
	if err != nil {
		return nil, err
	}
	return baseClasss, nil
}

// GetBaseClasssList 获取行政班列表
func GetBaseClasssList(page, pageSize int, filters ...interface{}) ([]*BaseClasss, int64) {
	offset := (page - 1) * pageSize
	list := make([]*BaseClasss, 0)
	query := orm.NewOrm().QueryTable("base_classs")
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
