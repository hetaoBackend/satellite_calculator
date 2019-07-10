// @APIVersion 1.0.0
// @Title beego satellite_calculate API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact hetao7@pku.edu.cn
// @TermsOfServiceUrl http://beego.me/
package routers

import (
	"satellite_calculator/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// 注释二处
		beego.NSNamespace("/satellite",
			beego.NSInclude(
				&controllers.SateController{},
			),
		),
		beego.NSNamespace("/station",
			beego.NSInclude(
				&controllers.StationController{},
			),
		),
		beego.NSNamespace("/city",
			beego.NSInclude(
				&controllers.CityController{},
			),
		),
		beego.NSNamespace("/calculate",
			beego.NSInclude(
				&controllers.CalculateController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
