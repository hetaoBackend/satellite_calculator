package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"satellite_calculator/models"
)

type StationController struct {
	beego.Controller
}

func (c *StationController) Get() {
	stationInfo := &models.StationInfo{Diameter: 0.5, TPower:3, TG:21.5, RGT:-1}
	c.Data["json"] = stationInfo
	c.ServeJSON()
}

func (c *StationController) Post() {
	stationInfo := models.SatelliteInfo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &stationInfo)
	id,_ := models.AddStation(stationInfo)
	data,_ := json.Marshal(map[string]int{"id": id})
	str := string(data[:])
	c.Data["json"] = str
	c.ServeJSON()
}


