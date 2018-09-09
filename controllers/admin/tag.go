package admin

import (
	"beegoblog/controllers"
	"beegoblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	"strconv"
	"time"
)

type TagController struct {
	controllers.BaseController
}

func (this *TagController) Get() {
	name := this.Input().Get("name")
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
	count, err := models.GetTagCount(name)
	if err != nil {
		beego.Error(err)
	}
	tags, err := models.GetTagList(name, pageNo, models.PageSize)
	if err != nil {
		beego.Error(err)
	}
	p := pagination.NewPaginator(this.Ctx.Request, models.PageSize, count)
	this.Data["Pager"] = p
	this.Data["List"] = tags
	this.Data["Name"] = name
	this.Data["Title"] = "标签"
	this.TplName = "admin/tag/index.html"
}

//添加和修改
func (this *TagController) Post() {
	strid := this.Input().Get("id")
	if len(strid) > 0 { //编辑
		id, err := strconv.ParseInt(strid, 10, 64)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "invalid params", true)
		}

		tag, err := models.GetTagOne(id)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "can not find this category", true)
		}
		tag.Name = this.Input().Get("name")
		tag.UpdatedAt = time.Now()
		err = models.Update(tag)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(500, err.Error(), true)
		}
	} else { //新增
		tag := &models.Tag{
			Name: this.Input().Get("name"),
		}
		err := models.Add(tag)
		if err != nil {
			beego.Error(err)
		}
	}
	this.ResponseJson(200, "success", true)
}

func (this *TagController) Delete() {
	beego.Info(this.GetInt64("id"))
	strid := this.Input().Get("id")
	if len(strid) > 0 {
		id, err := strconv.ParseInt(strid, 10, 64)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "invalid params", true)
		}
		tag := &models.Tag{
			Id: id,
		}
		err = models.Delete(tag)
		if err != nil {
			beego.Error(err)
			this.ResponseJson(400, "can not find this category", true)
		}
		this.ResponseJson(200, "success", true)
	} else {
		this.ResponseJson(400, "need id"+strid, true)
	}
}
