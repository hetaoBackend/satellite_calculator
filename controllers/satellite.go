package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"satellite_calculator/models"
	"strings"
)

type SateController struct {
	beego.Controller
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
	sateInfo.FreType = strings.ToLower(sateInfo.FreType)
	result, err := models.AddSate(sateInfo)
	if err != nil {
		logs.Error("satellite post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
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

// @Title get_satellites
// @Description Get satellites
// @Param    offset  query     int32    true        "the offset"
// @Param    count  query     int32    true        "the count"
// @Success 200 {object} []models.Satellite
// @Failure 403 body is empty
// @router /list/ [get]
func (c *SateController) GetSatellites() {
	offset, err := c.GetInt32("offset")
	if err != nil {
		logs.Error("invalid offset input")
		c.Abort("403")
	}
	count, err := c.GetInt32("count")
	if err != nil {
		logs.Error("invalid count input")
		c.Abort("403")
	}
	satesInfo, err := models.GetSatellites(offset, count)
	if err != nil {
		logs.Error("satellite get error, err: ", err)
		c.Data["json"] = make([]*models.Satellite, 0)
	} else {
		c.Data["json"] = satesInfo
	}
	c.ServeJSON()
}

// @Title delete satellite
// @Description Delete satellite
// @Param    name  query     string    true        "the name of satellite"
// @Param    fre_type  query     string    true        "the fre_type of satellite"
// @Success 200
// @Failure 403
// @router / [delete]
func (c *SateController) DeleteSatellite() {
	name := strings.ToUpper(c.GetString("name"))
	freType := strings.ToLower(c.GetString("fre_type"))
	err := models.DeleteSatellite(name, freType)
	if err != nil {
		logs.Error("delete satellite error, err: ", err)
		c.Abort("500")
	}
	c.Ctx.WriteString(fmt.Sprintf("delete %v %v success!", name, freType))
}
