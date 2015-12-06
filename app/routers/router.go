package routers

import (
	"github.com/GokulSrinivas/hackathon_pragyan/app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
