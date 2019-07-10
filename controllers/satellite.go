package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"satellite_calculator/models"
	"strings"
)

type SateController struct {
	beego.Controller
}

func (c *SateController) Get() {
	name := c.GetString("name")
	freType := strings.ToLower(c.GetString("fre_type"))
	sateInfo, err := models.GetSatellite(name, freType)
	if err != nil {
		logs.Error("satellite get error, err: ", err)
		c.Data["json"] = models.Satellite{}
	} else {
		c.Data["json"] = sateInfo
	}
	c.ServeJSON()
}

func (c *SateController) Post() {
	sateInfo := models.Satellite{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &sateInfo)
	if err != nil {
		logs.Error("satellite post error, err: ", err)
		c.Abort("500")
	}
	result, err := models.AddSate(sateInfo)
	if err != nil {
		logs.Error("satellite post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}
