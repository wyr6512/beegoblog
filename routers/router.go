package routers

import (
	"beegoblog/controllers"
	"beegoblog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/admin", &admin.HomeController{})
}
