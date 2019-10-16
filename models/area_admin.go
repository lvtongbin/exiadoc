/**********************************************
** @Des: 区域信息
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package models

import (
	"github.com/astaxie/beego/orm"
)

// BaseAreaAdmin 基本区域信息
type BaseAreaAdmin struct {
	ID    int    `json:"id" orm:"column(id);auto"`
	Admin string `json:"admin" orm:"column(admin)" description:"管理员账号"`
	Code  string `json:"code" orm:"column(code)" description:"编码"`
	Name  string `json:"name" orm:"column(name)" description:"区县名称"`
}

// TableName 表名
func (t *BaseAreaAdmin) TableName() string {
	return "base_area_admin"
}

func init() {
	orm.RegisterModel(new(BaseAreaAdmin))
}
