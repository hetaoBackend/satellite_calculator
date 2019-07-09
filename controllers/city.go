package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"satellite_calculator/models"
)

type CityController struct {
	beego.Controller
}

func (c *CityController) Get() {
	cityInfo := &models.City{Name:"北京", Longitude:116.5, Latitude:39.9}
	c.Data["json"] = cityInfo
	c.ServeJSON()
}

func (c *CityController) Post() {
	cityInfo := models.City{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &cityInfo)
	id,_ := models.AddCity(cityInfo)
	data,_ := json.Marshal(map[string]int64{"id": id})
	str := string(data[:])
	c.Data["json"] = str
	c.ServeJSON()
}

