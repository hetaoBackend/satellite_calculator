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
	cityInfo := &models.CityInfo{Name:"北京", Longtitude:116.5, Latitude:39.9}
	c.Data["json"] = cityInfo
	c.ServeJSON()
}

func (c *CityController) Post() {
	cityInfo := models.CityInfo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &cityInfo)
	id,_ := models.AddCity(cityInfo)
	data,_ := json.Marshal(map[string]int{"id": id})
	str := string(data[:])
	c.Data["json"] = str
	c.ServeJSON()
}

