package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"satellite_calculator/models"
)

type SateController struct {
	beego.Controller
}

func (c *SateController) Get() {
	name := c.GetString("name")
	fmt.Println(name)
	sateInfo, err := models.GetSatellite(name)
	if err != nil {
		fmt.Println("satellite get error, err: ", err)
		c.Data["json"] = models.Satellite{}
	} else {
		c.Data["json"] = sateInfo
	}
	c.ServeJSON()
}

func (c *SateController) Post() {
	sateInfo := models.Satellite{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &sateInfo)
	result, _ := models.AddSate(sateInfo)
	c.Data["json"] = result
	c.ServeJSON()
}
