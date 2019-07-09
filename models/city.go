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

func GetCity(name string) (*City, error) {
	if name == "" {
		return &City{}, nil
	}
	fmt.Println(name)
	p2 := City{Name: name}
	// 查询
	DB := orm.NewOrm()
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return &City{}, err
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return &City{}, err
	} else {
		fmt.Println(p2)
	}
	return &p2, nil
}

// ToDo: add city info to DB
func AddCity(cityInfo City) (int64, error) {
	DB := orm.NewOrm()
	id, err := DB.Insert(&cityInfo)
	if err != nil {
		fmt.Println("insert cityInfo error: ", err)
		return 0, err
	}
	return id, nil
}
