package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"satellite_calculator/models"
)

type StationController struct {
	beego.Controller
}

func (c *StationController) Get() {
	name := c.GetString("name")
	stationInfo, err := models.GetStation(name)
	if err != nil {
		fmt.Println("city get error, err: ", err)
		c.Data["json"] = models.City{}
	} else {
		c.Data["json"] = stationInfo
	}
	c.ServeJSON()
}

func (c *StationController) Post() {
	stationInfo := models.Station{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &stationInfo)
	result, err := models.AddStation(stationInfo)
	if err != nil {
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}


