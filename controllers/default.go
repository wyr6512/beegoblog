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

//全列表
func (this *MainController) Get() {
	page := this.Input().Get("p")
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
	params := models.ArticleParams{}
	articles, count, err := models.GetArticlePager(params, pageNo, models.PageSize)
	p := pagination.NewPaginator(this.Ctx.Request, models.PageSize, count)

	tags, err := models.GetTagList("", 1, 20)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Articles"] = articles
	this.Data["Pager"] = p
	this.Data["Tags"] = tags
	this.Data["Title"] = "beego blog"
	this.TplName = "index.tpl"
}

//分类文章列表
func (this *MainController) GetCateArticles() {
	id := this.Ctx.Input.Param(":id")
	cateId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	page := this.Input().Get("p")
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
	params := models.ArticleParams{
		CategoryId: cateId,
	}
	articles, count, err := models.GetArticlePager(params, pageNo, models.PageSize)
	p := pagination.NewPaginator(this.Ctx.Request, models.PageSize, count)
	tags, err := models.GetTagList("", 1, 20)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Type"] = "home"
	this.Data["Articles"] = articles
	this.Data["Pager"] = p
	this.Data["Tags"] = tags
	this.Data["Title"] = "category articles"
	this.TplName = "index.tpl"
}

//标签列表文章
func (this *MainController) GetTagArticles() {
	name := this.Ctx.Input.Param(":name")
	tag, err := models.GetTagsByNames([]string{name})
	var tagId int64
	if err != nil {
		tagId = 0
		beego.Error(err)
	}
	tagId = tag[0].Id
	if err != nil {
		beego.Error(err)
	}
	page := this.Input().Get("p")
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
	articleTags, err := models.GetArticleTagList(tagId, pageNo, models.PageSize)
	if err != nil {
		beego.Error(err)
	}
	var articleIds []int64
	for _, v := range articleTags {
		articleIds = append(articleIds, v.ArticleId)
	}
	articles := []*models.Article{}
	var count int64
	p := pagination.NewPaginator(this.Ctx.Request, models.PageSize, 0)
	if len(articleIds) > 0 {
		params := models.ArticleParams{
			Ids: articleIds,
		}
		articles, count, err = models.GetArticlePager(params, pageNo, models.PageSize)
		p = pagination.NewPaginator(this.Ctx.Request, models.PageSize, count)
	}
	tags, err := models.GetTagList("", 1, 20)
	if err != nil {
		beego.Error(err)
	}
	this.Data["IsTag"] = true
	this.Data["Tag"] = tag[0]
	this.Data["Articles"] = articles
	this.Data["Pager"] = p
	this.Data["Tags"] = tags
	this.Data["Title"] = "tag articles"
	this.TplName = "index.tpl"
}

//文章详情
func (this *MainController) GetArticleDetail() {
	id := this.Ctx.Input.Param(":id")
	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	article, err := models.GetArticleOne(articleId)
	if err != nil {
		beego.Error(err)
	}
	content, err := models.GetContentByArticleId(articleId)
	if err != nil {
		beego.Error(err)
	}
	tags, err := models.GetTagList("", 1, 20)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Article"] = article
	this.Data["Content"] = content
	this.Data["Tags"] = tags
	this.Data["Title"] = "article detail"
	this.TplName = "article_detail.html"
}
