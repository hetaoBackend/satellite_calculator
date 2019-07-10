package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"satellite_calculator/models"
	"strings"
)

type StationController struct {
	beego.Controller
}

func (c *StationController) Get() {
	name := c.GetString("name")
	freType := strings.ToLower(c.GetString("fre_type"))
	stationInfo, err := models.GetStation(name, freType)
	if err != nil {
		logs.Error("station get error, err: ", err)
		c.Data["json"] = models.Station{}
	} else {
		c.Data["json"] = stationInfo
	}
	c.ServeJSON()
}

func (c *StationController) Post() {
	stationInfo := models.Station{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &stationInfo)
	if err != nil {
		logs.Error("station post error, err: ", err)
		c.Abort("500")
	}
	result, err := models.AddStation(stationInfo)
	if err != nil {
		logs.Error("station post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}


