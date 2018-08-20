package admin

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["Title"] = "首页"
	this.TplName = "admin/index.html"
}
