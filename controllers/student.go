/**********************************************
** @Des: 学生管理
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package controllers

import (
	"fmt"
	"trism/models"
	"trism/utils"

	"github.com/astaxie/beego"
)

// BaseStudentController operations for BaseStudent
type BaseStudentController struct {
	beego.Controller
}

// AddStudents is ...
// @Title AddStudents
// @Description 批量添加随机学生
// @router /addStudents [get]
func (c *BaseStudentController) AddStudents() {
	for i := 0; i < 10000; i++ {
		id := utils.GetRandomString(18)
		name := utils.GetNameRandom()
		school, nianji, class := utils.GetSchoolRandom()
		student := &models.BaseStudent{
			Batch:     nianji,
			Idcard:    id,
			Name:      name,
			School:    school,
			BaseClass: class,
		}
		if _, err := models.AddBaseStudent(student); err == nil {
			fmt.Println("success")
		}
	}
	c.Data["json"] = map[string]int64{"code": 0}
	c.ServeJSON()
}
