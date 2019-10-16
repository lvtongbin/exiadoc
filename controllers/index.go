package controllers

import (
	"trism/models"
	"trism/utils"

	"github.com/astaxie/beego"
)

// IndexController is Controller of index
type IndexController struct {
	beego.Controller
}

// Index 登录入口
func (index *IndexController) Index() {
	// 获取年级
	batchList, _ := models.GetBatchList(1, 99)

	index.Data["gradeList"] = batchList
	index.Data["token"] = utils.GetCodeString(12)
	index.TplName = "index/index.html"
}
