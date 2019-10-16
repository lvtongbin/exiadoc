package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

const (
	// MessageOK 正确消息返回
	MessageOK = 0
	// MessageError 错误消息返回
	MessageError = -1
)

// BaseController is controller of base.
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

// Prepare 前期准备
func (baseController *BaseController) Prepare() {
	controllerName, actionName := baseController.GetControllerAndAction()
	baseController.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	baseController.actionName = strings.ToLower(actionName)
}

// ajaxMsg ajax返回
func (baseController *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["message"] = msg
	baseController.Data["json"] = out
	baseController.ServeJSON()
	baseController.StopRun()
}

// ajaxList ajax返回 列表
func (baseController *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["message"] = msg
	out["count"] = count
	out["data"] = data
	baseController.Data["json"] = out
	baseController.ServeJSON()
	baseController.StopRun()
}
