package routers

import (
	"satellite_calculator/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/satellite", &controllers.SateController{})
	beego.Router("/station", &controllers.StationController{})
	beego.Router("/city", &controllers.CityController{})
	beego.Router("/calculate", &controllers.CalculateController{})
}
