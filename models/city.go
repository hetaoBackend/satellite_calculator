package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
	orm.RegisterModel(new(City))
}

type City struct {
	Id int `orm:"pk"`
	Name string
	Longitude float32
	Latitude float32
}

// ToDo: add city info to DB
func AddCity(cityInfo City) (int64, error) {
	DB := orm.NewOrm()
	fmt.Println(cityInfo)
	id, err := DB.Insert(&cityInfo)
	if err != nil {
		fmt.Println("insert cityInfo error: ", err)
		return 0, err
	}
	return id, nil
}
