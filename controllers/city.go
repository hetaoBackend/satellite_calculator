package controllers

import (
	"encoding/json"
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

