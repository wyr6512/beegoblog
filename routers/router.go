package routers

import (
	"beegoblog/controllers"
	"beegoblog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/category/:name", &controllers.MainController{}, "*:GetCateArticles")
	beego.Router("/tag/:name", &controllers.MainController{}, "*:GetTagArticles")
	beego.Router("/article/:id", &controllers.MainController{}, "*:GetArticleDetail")

	beego.Router("/login", &admin.LoginController{})
	beego.Router("/admin", &admin.HomeController{})
	beego.Router("/admin/category/all", &admin.CategoryController{}, "*:GetAll")
	beego.Router("/admin/category/delete", &admin.CategoryController{}, "*:Delete")
	beego.Router("/admin/category", &admin.CategoryController{})
	beego.Router("/admin/article", &admin.ArticleController{})
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/upload", &admin.UploadController{}, "*:Upload")
	beego.Router("/admin/tag/delete", &admin.TagController{}, "*:Delete")
	beego.Router("/admin/tag", &admin.TagController{})
}
