package admin

import (
	"beegoblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"

	"strconv"
)

type CategoryController struct {
	beego.Controller
}

//列表
func (this *CategoryController) Get() {
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
	count, err := models.GetCategoryCount(name)
	if err != nil {
		beego.Error(err)
	}
	cates, err := models.GetCategoryList(name, pageNo, models.PageSize)
	if err != nil {
		beego.Error(err)
	}
	p := pagination.NewPaginator(this.Ctx.Request, models.PageSize, count)
	this.Data["Pager"] = p
	this.Data["List"] = cates
	this.Data["Website"] = "beego.me"
	this.Data["Title"] = "分类"
	this.TplName = "admin/category/index.html"
}

//添加和修改
func (this *CategoryController) Post() {
	strid := this.Input().Get("id")
	if len(strid) > 0 { //编辑
		id, err := strconv.ParseInt(strid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		cate := &models.Category{
			Id:   id,
			Name: this.Input().Get("name"),
		}
		err = models.UpdateCategory(cate)
		if err != nil {
			beego.Error(err)
		}
	} else { //新增
		cate := &models.Category{
			Name: this.Input().Get("name"),
		}
		err := models.AddCategory(cate)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["json"] = map[string]interface{}{"code": 200, "msg": "success"}
	this.ServeJSON()
}
