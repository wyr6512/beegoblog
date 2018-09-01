package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ResponseJson(code int, msg interface{}, stop bool) {
	this.Ctx.ResponseWriter.WriteHeader(code)
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	this.ServeJSON()
	if stop {
		this.StopRun()
	}
}
