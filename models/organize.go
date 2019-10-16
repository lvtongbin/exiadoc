/**********************************************
** @Des: 学校基础信息
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package models

import "github.com/astaxie/beego/orm"

// BaseOrganize 学校基础信息表
type BaseOrganize struct {
	ID       int    `json:"id" orm:"column(id)"`
	Code     string `json:"code" orm:"column(code)" description:"编码"`
	AreaCode int    `json:"areacode" orm:"column(area_code)" description:"区县编码"`
	Name     string `json:"name" orm:"column(name)" description:"学校名称"`
}

// TableName 表名
func (t *BaseOrganize) TableName() string {
	return "base_organize"
}

func init() {
	orm.RegisterModel(new(BaseOrganize))
}

// GetBaseOrganizeByName 获取学校信息
func GetBaseOrganizeByName(name string) (*BaseOrganize, error) {
	baseOrganize := new(BaseOrganize)
	err := orm.NewOrm().QueryTable(baseOrganize).Filter("Name", name).One(baseOrganize)
	if err != nil {
		return nil, err
	}
	return baseOrganize, nil
}

// GetSchoolList 获取学校列表
func GetSchoolList(page, pageSize int, filters ...interface{}) ([]*BaseOrganize, int64) {
	offset := (page - 1) * pageSize
	list := make([]*BaseOrganize, 0)
	query := orm.NewOrm().QueryTable("base_organize")
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
