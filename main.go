package main

import (
	"beegoblog/models"
	_ "beegoblog/routers"
	"beegoblog/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"net/http"
	"strings"
)

func main() {
	beego.AddFuncMap("processTags", processTags)
	beego.AddFuncMap("getCateName", getCateName)
	//后台鉴权
	beego.InsertFilter("/admin/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie, err := ctx.Request.Cookie("Authorization")
		beego.Error(cookie)
		if err != nil || !util.CheckToken(cookie.Value) {
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", 302)
		}
	})
	beego.Run()
}

func processTags(in string) (out string) {
	out = strings.Trim(strings.Replace(in, "$#", ",", -1), "#$")
	return
}

func getCateName(in int64) (out string) {
	cate, err := models.GetCategoryOne(in)
	if err != nil {
		out = ""
	} else {
		out = cate.Name
	}
	return
}
