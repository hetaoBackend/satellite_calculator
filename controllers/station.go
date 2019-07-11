package controllers

import (
	"encoding/json"
	"fmt"
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
	stationInfo.FreType = strings.ToLower(stationInfo.FreType)
	result, err := models.AddStation(stationInfo)
	if err != nil {
		logs.Error("station post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Title get_stations
// @Description Get stations
// @Param    offset  query     int32    true        "the offset"
// @Param    count  query     int32    true        "the count"
// @Success 200 {object} []models.Station
// @Failure 403 body is empty
// @router /list/ [get]
func (c *StationController) GetStations() {
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
	stationsInfo, err := models.GetStations(offset, count)
	if err != nil {
		logs.Error("station get error, err: ", err)
		c.Data["json"] = make([]*models.Station, 0)
	} else {
		c.Data["json"] = stationsInfo
	}
	c.ServeJSON()
}

// @Title delete station
// @Description Delete station
// @Param    name  query     string    true        "the name of station"
// @Param    fre_type  query     string    true        "the fre_type of station"
// @Success 200
// @Failure 403
// @router / [delete]
func (c *StationController) DeleteStation() {
	name := strings.ToUpper(c.GetString("name"))
	freType := strings.ToLower(c.GetString("fre_type"))
	err := models.DeleteStation(name, freType)
	if err != nil {
		logs.Error("delete station error, err: ", err)
		c.Abort("500")
	}
	c.Ctx.WriteString(fmt.Sprintf("delete %v %v success!", name, freType))
}

