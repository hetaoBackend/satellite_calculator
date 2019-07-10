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

// @Title get
// @Description Get station
// @Param    name  query     string    true        "the name of satellite"
// @Param    fre_type  query     string    true        "the frequency type of satellite"
// @Success 200 {object} models.Satellite
// @Failure 403 body is empty
// @router / [get]
func (c *SateController) Get() {
	name := strings.ToUpper(c.GetString("name"))
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

// @Title post
// @Description create/modify satellite
// @Param   body        body    models.Satellite   true        "The satellite content"
// @Success 200 {object} models.Satellite
// @Failure 403 body is empty
// @router / [post]
func (c *SateController) Post() {
	sateInfo := models.Satellite{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &sateInfo)
	if err != nil {
		logs.Error("satellite post error, err: ", err)
		c.Abort("500")
	}
	sateInfo.Name = strings.ToUpper(sateInfo.Name)
	result, err := models.AddSate(sateInfo)
	if err != nil {
		logs.Error("satellite post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}
