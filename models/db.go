package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	maxIdle := 30
	maxConn := 30
	// 创建连接池
	err := orm.RegisterDataBase("default", "mysql", "satellite_lab_pku:pku_2935@tcp(127.0.0.1:13306)/link_budget?charset=utf8", maxIdle, maxConn)

	if err != nil {
		fmt.Println("connect mysql err : ", err)
	}
	// 需要在init中注册定义的model
	orm.RegisterModel(new(City), new(Station), new(Satellite))
}
