package controllers

import (
	"beegoblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	"strconv"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	page := this.Input().Get("p")
	params := models.ArticleParams{}
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

	tags, err := models.GetTagList("", pageNo, 20)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Articles"] = articles
	this.Data["Pager"] = p
	this.Data["Tags"] = tags
	this.TplName = "index.tpl"
}
