// @APIVersion 1.0.1
// @Title satellite_calculator API
// @Description An useful tool to do the satellite link budget calculate
// @Contact hetao7@pku.edu.cn
// @TermsOfServiceUrl http://beego.me/
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"satellite_calculator/controllers"
	"strings"
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
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	if strings.Index(ctx.Request.URL.Path, "v1/") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)
}
