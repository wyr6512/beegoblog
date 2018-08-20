package admin

import (
	"beegoblog/models"
	"github.com/astaxie/beego"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	name:=this.Input().Get("name")
	page:=this.Input().Get("page")
	var pageNo int64
	if len(page) == 0{
		pageNo = 1
	}
	pageNo, err:=strconv.ParseInt(page,10,64)
	if err!=nil{
		beego.Error(err)
	}

	pager, err:=models.GetCategoryList(name,pageNo,models.PageSize)
	if err!=nil{
		beego.Error(err)
	}
	this.Data["Pager"] = pager
	this.Data["Website"] = "beego.me"
	this.Data["Title"] = "分类"
	this.TplName = "admin/category/index.html"
}
