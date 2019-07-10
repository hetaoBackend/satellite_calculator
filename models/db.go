package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	maxIdle := 30
	maxConn := 30
	// 创建连接池
	databaseURL := "127.0.0.1:13306"
	if beego.BConfig.RunMode == "docker" {
		databaseURL = "link_budget_db:3306"
	}
	err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("satellite_lab_pku:pku_2935@tcp(%s)/link_budget?charset=utf8", databaseURL), maxIdle, maxConn)

	if err != nil {
		fmt.Println("connect mysql err : ", err)
	}
	// 需要在init中注册定义的model
	orm.RegisterModel(new(City), new(Station), new(Satellite))
}
