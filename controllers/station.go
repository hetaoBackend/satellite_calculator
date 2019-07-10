package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"satellite_calculator/models"
	"strings"
)

type StationController struct {
	beego.Controller
}

// @Title get
// @Description Get Station
// @Param    name  query     string    true        "the name of station"
// @Param    fre_type  query     string    true        "the frequency type of station"
// @Success 200 {object} models.Station
// @Failure 403 body is empty
// @router / [get]
func (c *StationController) Get() {
	name := strings.ToUpper(c.GetString("name"))
	freType := strings.ToLower(c.GetString("fre_type"))
	stationInfo, err := models.GetStation(name, freType)
	if err != nil {
		logs.Error("station get error, err: ", err)
		c.Data["json"] = models.Station{}
	} else {
		c.Data["json"] = stationInfo
	}
	c.ServeJSON()
}

// @Title post
// @Description create/modify station
// @Param   body        body    models.Station   true        "The station content"
// @Success 200 {object} models.Station
// @Failure 403 body is empty
// @router / [post]
func (c *StationController) Post() {
	stationInfo := models.Station{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &stationInfo)
	if err != nil {
		logs.Error("station post error, err: ", err)
		c.Abort("500")
	}
	stationInfo.Name = strings.ToUpper(stationInfo.Name)
	result, err := models.AddStation(stationInfo)
	if err != nil {
		logs.Error("station post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}


