package admin

import (
	"github.com/astaxie/beego"
)

type CategoryController struct{
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["Website"] = "beego.me"
	this.TplName = "admin/category/index.html"
}