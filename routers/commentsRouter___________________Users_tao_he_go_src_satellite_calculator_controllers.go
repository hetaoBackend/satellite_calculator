package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["satellite_calculator/controllers:CalculateController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:CalculateController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:CityController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:CityController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:CityController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:CityController"],
        beego.ControllerComments{
            Method: "DeleteCity",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:CityController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:CityController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:CityController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetCities",
            Router: `/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:SateController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:SateController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:SateController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:SateController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:SateController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:SateController"],
        beego.ControllerComments{
            Method: "DeleteSatellite",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:SateController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:SateController"],
        beego.ControllerComments{
            Method: "GetSatellites",
            Router: `/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:StationController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:StationController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:StationController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:StationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:StationController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:StationController"],
        beego.ControllerComments{
            Method: "DeleteStation",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["satellite_calculator/controllers:StationController"] = append(beego.GlobalControllerRouter["satellite_calculator/controllers:StationController"],
        beego.ControllerComments{
            Method: "GetStations",
            Router: `/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
