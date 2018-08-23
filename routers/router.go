package routers

import (
	"beegoblog/controllers"
	"beegoblog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.HomeController{})
	beego.Router("/admin/category/delete", &admin.CategoryController{}, "*:Delete")
	beego.Router("/admin/category", &admin.CategoryController{})
}
