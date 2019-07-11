package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"satellite_calculator/models"
	"strings"
)

type CityController struct {
	beego.Controller
}

// @Title get
// @Description Get city
// @Param    name  query     string    true        "the name of city"
// @Success 200 {object} models.City
// @Failure 403 body is empty
// @router / [get]
func (c *CityController) Get() {
	name := strings.ToUpper(c.GetString("name"))
	cityInfo, err := models.GetCity(name)
	if err != nil {
		logs.Error("city get error, err: ", err)
		c.Data["json"] = models.City{}
	} else {
		c.Data["json"] = cityInfo
	}
	c.ServeJSON()
}

// @Title get_cities
// @Description Get cities
// @Param    offset  query     int32    true        "the offset"
// @Param    count  query     int32    true        "the count"
// @Success 200 {object} []models.City
// @Failure 403 body is empty
// @router /list/ [get]
func (c *CityController) GetCities() {
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
	citiesInfo, err := models.GetCities(offset, count)
	if err != nil {
		logs.Error("city get error, err: ", err)
		c.Data["json"] = make([]*models.City, 0)
	} else {
		c.Data["json"] = citiesInfo
	}
	c.ServeJSON()
}

// @Title delete city
// @Description Delete city
// @Param    name  query     string    true        "the name of city"
// @Success 200
// @Failure 403
// @router / [delete]
func (c *CityController) DeleteCity() {
	name := strings.ToUpper(c.GetString("name"))
	err := models.DeleteCity(name)
	if err != nil {
		logs.Error("delete city error, err: ", err)
		c.Abort("500")
	}
	c.Ctx.WriteString(fmt.Sprintf("delete %v success!", name))
}

// @Title post
// @Description create/modify city
// @Param   body        body    models.City   true        "The city content"
// @Success 200 {object} models.City
// @Failure 403 body is empty
// @router / [post]
func (c *CityController) Post() {
	cityInfo := models.City{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cityInfo)
	if err != nil {
		logs.Error("city post error, err: ", err)
		c.Abort("500")
	}
	cityInfo.Name = strings.ToUpper(cityInfo.Name)
	result, err := models.AddCity(cityInfo)
	if err != nil {
		logs.Error("city post error, err: ", err)
		c.Abort("403")
	}
	c.Data["json"] = result
	c.ServeJSON()
}

