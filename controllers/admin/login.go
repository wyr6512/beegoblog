package admin

import (
	"beegoblog/controllers"
	"beegoblog/util"
	"github.com/astaxie/beego"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "yes"
	if isExit {
		this.Ctx.SetCookie("Authorization", "", -100, "/")
		this.Ctx.SetCookie("uname", "", -100, "/")
		this.Redirect("/login", 302)
		return
	}
	this.Data["Title"] = "登录"
	this.TplName = "admin/login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		token := util.GenToken()
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("Authorization", token, maxAge, "/")
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Redirect("/admin", 302)
		return
	}
	this.Redirect("/login", 302)
	return
}
