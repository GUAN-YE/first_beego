package routers

import (
	"ceshi_go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/photos",&controllers.FormContrller{})
}
