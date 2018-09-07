package routers

import (
	"beegoblog/controllers"
	"beegoblog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.HomeController{})
	beego.Router("/admin/category/all", &admin.CategoryController{}, "*:GetAll")
	beego.Router("/admin/category/delete", &admin.CategoryController{}, "*:Delete")
	beego.Router("/admin/category", &admin.CategoryController{})
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article", &admin.ArticleController{})
	beego.Router("/admin/upload", &admin.UploadController{}, "*:Upload")
}
