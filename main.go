package main

import (
	"github.com/astaxie/beego/logs"
	_ "satellite_calculator/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	logs.SetLogger("file", `{"filename":"service.log"}`)
	beego.Run()
}

