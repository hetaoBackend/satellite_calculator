package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"satellite_calculator/models"
)

type CityController struct {
	beego.Controller
}

func (c *CityController) Get() {
	name := c.GetString("name")
	cityInfo, err := models.GetCity(name)
	if err != nil {
		logs.Error("city get error, err: ", err)
		c.Data["json"] = models.City{}
	} else {
		c.Data["json"] = cityInfo
	}
	c.ServeJSON()
}

func (c *CityController) Post() {
	cityInfo := models.City{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cityInfo)
	if err != nil {
		logs.Error("city post error, err: ", err)
		c.Abort("500")
	}
	result, err := models.AddCity(cityInfo)
	if err != nil {
		logs.Error("city post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}

