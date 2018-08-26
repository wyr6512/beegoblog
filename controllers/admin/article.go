package admin

import (
	"beegoblog/controllers"
	"beegoblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	"strconv"
	//"time"
	// "fmt"
)

type ArticleController struct {
	controllers.BaseController
}

func (this *ArticleController) Add() {
	this.Data["Title"] = "文章-添加"
	this.TplName = "admin/article/add.html"
}

//列表
func (this *ArticleController) Get() {
	page := this.Input().Get("p")
	strStatus := this.Input().Get("status")
	var status int
	if len(strStatus) == 0 {
		status = 0
	} else {
		var err error
		status, err = strconv.Atoi(strStatus)
		if err != nil {
			beego.Error(err)
		}
	}
	params := models.ArticleParams{
		Title:    this.Input().Get("title"),
		Status:   status,
		FromTime: this.Input().Get("from"),
		EndTime:  this.Input().Get("end"),
		Tag:      this.Input().Get("tags"),
	}
	var pageNo int64
	if len(page) == 0 {
		pageNo = 1
	} else {
		var err error
		pageNo, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	count, err := models.GetArticleCount(params)
	if err != nil {
		beego.Error(err)
	}
	articles, err := models.GetArticleList(params, pageNo, models.PageSize)
	if err != nil {
		beego.Error(err)
	}
	p := pagination.NewPaginator(this.Ctx.Request, models.PageSize, count)
	this.Data["Pager"] = p
	this.Data["List"] = articles

	this.Data["Title"] = "文章列表"
	this.TplName = "admin/article/index.html"
}
