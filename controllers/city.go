package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"satellite_calculator/models"
)

type CityController struct {
	beego.Controller
}

func (c *CityController) Get() {
	name := c.GetString("name")
	cityInfo, err := models.GetCity(name)
	if err != nil {
		fmt.Println("city get error, err: ", err)
		c.Data["json"] = models.City{}
	} else {
		c.Data["json"] = cityInfo
	}
	c.ServeJSON()
}

func (c *CityController) Post() {
	cityInfo := models.City{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &cityInfo)
	result, err := models.AddCity(cityInfo)
	if err != nil {
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}

