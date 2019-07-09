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
	id,_ := models.AddCity(cityInfo)
	data,_ := json.Marshal(map[string]int64{"id": id})
	str := string(data[:])
	c.Data["json"] = str
	c.ServeJSON()
}

