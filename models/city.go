package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

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
func AddCity(cityInfo City) (City, error) {
	DB := orm.NewOrm()
	id, err := DB.Insert(&cityInfo)
	if err != nil {
		fmt.Println("insert cityInfo error: ", err)
		return City{}, err
	}
	cityInfo.Id = int(id)
	return cityInfo, nil
}
