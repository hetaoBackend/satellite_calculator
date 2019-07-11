package main

import (
	"github.com/astaxie/beego/logs"
	"os"
	_ "satellite_calculator/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "docker" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	err := os.MkdirAll("logs/", 0775)
	if err != nil {
		logs.Error("mkdir log error: ", err)
		return
	}
	err = logs.SetLogger("file", `{"filename":"logs/service.log"}`)
	if err != nil {
		logs.Error("set logger error: ", err)
		return
	}
	beego.Run()
}

