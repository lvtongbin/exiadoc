/**********************************************
** @Des: 年级信息
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package models

import (
	"github.com/astaxie/beego/orm"
)

// ItemsBatch 年级信息表
type ItemsBatch struct {
	ID   int    `json:"id" orm:"column(id);auto"`
	Code string `json:"code" orm:"column(code)" description:"编码"`
	Name string `json:"name" orm:"column(name)" description:"年级名"`
}

// TableName 表名
func (t *ItemsBatch) TableName() string {
	return "items_batch"
}

func init() {
	orm.RegisterModel(new(ItemsBatch))
}

// AddItemsBatch insert a new ItemsBatch into database and returns
// last inserted Id on success.
func AddItemsBatch(m *ItemsBatch) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetItemsBatch 获取年级信息
func GetItemsBatch(name string) (*ItemsBatch, error) {
	batch := new(ItemsBatch)
	err := orm.NewOrm().QueryTable(batch).Filter("name", name).One(batch)
	return batch, err
}

// GetBatchList 获取年级列表
func GetBatchList(page, pageSize int, filters ...interface{}) ([]*ItemsBatch, int64) {
	offset := (page - 1) * pageSize
	list := make([]*ItemsBatch, 0)
	query := orm.NewOrm().QueryTable("items_batch")
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
