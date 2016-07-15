package routers

import (
	"github.com/golang-snippet/beegoweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.ErrorController(&controllers.ErrorController{})
    beego.Router("/", &controllers.MainController{})
    beego.Router("/bla", &controllers.BlaController{})
		beego.Router("/tab", &controllers.TabController{})
}
