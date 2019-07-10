package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"satellite_calculator/models"
)

type CalculateController struct {
	beego.Controller
}


func (c *CalculateController) Post() {
	calculateInfo := models.CalculateRequest{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &calculateInfo)
	if err != nil {
		logs.Error("calculate post error, err: ", err)
		c.Abort("500")
	}
	result, err := models.Calculate(calculateInfo)
	if err != nil {
		logs.Error("calculate post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}