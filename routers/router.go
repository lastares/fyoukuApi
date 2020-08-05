package routers

import (
	"fyoukuApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UserController{})
    beego.Include(&controllers.VideoController{})
    beego.Include(&controllers.BaseController{})
    beego.Include(&controllers.CommonController{})
    beego.Include(&controllers.TopController{})
    beego.Include(&controllers.BarrageController{})
    beego.Include(&controllers.WsController{})
}
