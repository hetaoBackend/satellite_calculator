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

// @Title post
// @Description calculate link_budget
// @Param   body        body    models.CalculateRequest   true        "The CalculateRequest content"
// @Success 200 {object} models.CalculateResponse
// @Failure 403 body is empty
// @router / [post]
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