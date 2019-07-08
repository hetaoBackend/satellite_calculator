package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"satellite_calculator/models"
)

type SateController struct {
	beego.Controller
}

func (c *SateController) Get() {
	myStruct := &models.SatelliteInfo{Name: "ST-2号", GT: 1, SFD: -90, EIRP: 46, Longtitude: 98.2}
	c.Data["json"] = myStruct
	c.ServeJSON()
}

func (c *SateController) Post() {
	ob := models.SatelliteInfo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	id, _ := models.AddSate(ob)
	data, _ := json.Marshal(map[string]int{"id": id})
	str := string(data[:])
	c.Data["json"] = str
	c.ServeJSON()
}
