package admin

import (
	"beegoblog/controllers"
	"beegoblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	"strconv"
	"strings"
	"time"
	// "fmt"
)

type ArticleController struct {
	controllers.BaseController
}

func (this *ArticleController) Add() {
	this.Data["Title"] = "文章-添加"
	this.TplName = "admin/article/add.html"
}

func (this *ArticleController) Edit() {
	strId := this.Input().Get("id")
	var id int64
	if len(strId) == 0 {
		id = 0
	} else {
		var err error
		id, err = strconv.ParseInt(strId, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		article, err := models.GetArticleOne(id)
		if err != nil {
			beego.Error(err)
		}
		content, err := models.GetContentByArticleId(article.Id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Article"] = article
		this.Data["Content"] = content
	}
	this.Data["Title"] = "文章-编辑"
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
	title := this.Input().Get("title")
	params := models.ArticleParams{
		Title:    title,
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
	this.Data["Query"] = title
	this.Data["Title"] = "文章列表"
	this.TplName = "admin/article/index.html"
}

func (this *ArticleController) Post() {
	strId := this.Input().Get("id")
	categoryId, err := strconv.ParseInt(this.Input().Get("category"), 10, 64)
	if err != nil {
		beego.Error(err)
		this.ResponseJson(400, err.Error(), true)
	}
	if len(strId) == 0 { //新增
		strTags := this.Input().Get("tags")
		tags := strTags
		if len(strTags) > 0 {
			tags = strings.Replace(strTags, ",", "#$", -1)
			tags = "#" + tags + "$"
		}
		article := &models.Article{
			Title:      this.Input().Get("title"),
			Abstract:   this.Input().Get("abstract"),
			CategoryId: categoryId,
			Tags:       tags,
		}
		err = models.AddArticle(article, this.Input().Get("content"))
		if err != nil {
			beego.Error(err)
			this.ResponseJson(500, err.Error(), true)
		}
	} else { //编辑
		id, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, err.Error(), true)
		}
		article, err := models.GetArticleOne(id)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(500, err.Error(), true)
		}
		content, err := models.GetContentByArticleId(article.Id)
		article.UpdatedAt = time.Now()
		content.UpdatedAt = time.Now()
		err = models.UpdateArticle(article, content)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(500, err.Error(), true)
		}
	}
	this.ResponseJson(200, "success", true)
}

func (this *ArticleController) Delete() {
	beego.Info(this.GetInt64("id"))
	strid := this.Input().Get("id")
	if len(strid) > 0 {
		id, err := strconv.ParseInt(strid, 10, 64)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "invalid params", true)
		}
		article := &models.Article{
			Id: id,
		}
		err = models.DeleteArticle(article)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "can not find this article", true)
		}
		this.ResponseJson(200, "success", true)
	} else {
		this.ResponseJson(400, "need id"+strid, true)
	}
}
