package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Satellite struct {
	Id int `orm:"pk"`
	Name string
	GT float32
	SFD float32
	EIRP float32
	Longitude float32
}

func GetSatellite(name string) (*Satellite, error) {
	if name == "" {
		return &Satellite{}, nil
	}
	p2 := Satellite{Name: name}
	// 查询
	DB := orm.NewOrm()
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return &Satellite{}, err
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return &Satellite{}, err
	} else {
		fmt.Println(p2)
	}
	return &p2, nil
}

// ToDo: add satellite info to DB
func AddSate(sateInfo Satellite) (Satellite, error) {
	DB := orm.NewOrm()
	id, err := DB.Insert(&sateInfo)
	if err != nil {
		fmt.Println("insert sateInfo error: ", err)
		return Satellite{}, err
	}
	sateInfo.Id = int(id)
	return sateInfo, nil
}