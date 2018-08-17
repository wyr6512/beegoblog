package admin

import (
	"github.com/astaxie/beego"
)

type HomeController struct{
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "admin/index.html"
}