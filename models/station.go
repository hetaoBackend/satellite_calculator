package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Station struct {
	Id int `orm:"pk"`
	Name string
	Diameter float32
	TPower float32
	TG float32
	RGT float32
}

func GetStation(name string) (*Station, error) {
	if name == "" {
		return &Station{}, nil
	}
	fmt.Println(name)
	p2 := Station{Name: name}
	// 查询
	DB := orm.NewOrm()
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return &Station{}, err
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return &Station{}, err
	} else {
		fmt.Println(p2)
	}
	return &p2, nil
}

// ToDo: add station info to DB
func AddStation(stationInfo Station) (Station, error) {
	DB := orm.NewOrm()
	id, err := DB.Insert(&stationInfo)
	if err != nil {
		fmt.Println("insert cityInfo error: ", err)
		return Station{}, err
	}
	stationInfo.Id = int(id)
	return stationInfo, nil
}
