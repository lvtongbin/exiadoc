/**********************************************
** @Des: 班级管理
** @Author: lvtongbin
** @Date:   2019-08-17
** @Last Modified by:   lvtongbin
** @Last Modified time: 2019-09-09 18:50:41
***********************************************/

package controllers

import (
	"strconv"
	"strings"
	"time"
	"trism/models"

	"github.com/astaxie/beego"
)

// ClasssController operations for Classs
type ClasssController struct {
	beego.Controller
}

// InitBaseClass is ...
// @Title InitBaseClass
// @Description 初始化行政班级
// @router /initBaseClass [get]
func (c *ClasssController) InitBaseClass() {
	classbase := []string{"级一班", "级二班", "级三班", "级四班"}
	filters := make([]interface{}, 0)
	// 获取学校列表
	schoollist, count := models.GetSchoolList(1, 99, filters...)
	if count <= 0 {
		c.Data["json"] = map[string]int64{"code": 1001}
		c.ServeJSON()
	}
	cpschool := 0
	chschool := 0
	for _, school := range schoollist {
		codebase := 1000000
		if strings.Contains(school.Name, "中学") {
			codebase = 2000000
			chschool++
		} else {
			cpschool++
		}
		if codebase == 1000000 {
			for times := 0; times < 5; times++ {
				for k, class := range classbase {
					name := strconv.Itoa(15+times) + class
					code := strconv.Itoa(codebase + 10000*cpschool + 10*(times+1) + k + 1)
					batch := "小学" + strconv.Itoa(15+times) + "级"
					v := &models.BaseClasss{
						Name:   name,
						Code:   code,
						School: school.Name,
						Batch:  batch,
					}
					models.AddBaseClasss(v)
				}
			}
		} else {
			for times := 0; times < 3; times++ {
				for k, class := range classbase {
					name := strconv.Itoa(17+times) + class
					code := strconv.Itoa(codebase + 10000*chschool + 10*(times+1) + k + 1)
					batch := "中学" + strconv.Itoa(17+times) + "级"
					v := &models.BaseClasss{
						Name:   name,
						Code:   code,
						School: school.Name,
						Batch:  batch,
					}
					models.AddBaseClasss(v)
				}
			}
		}
	}
	c.Data["json"] = map[string]int64{"timestamp": time.Now().Unix()}
	c.ServeJSON()
}
