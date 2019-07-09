package models

import (
	"errors"
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
	if stationInfo.Diameter == 0 && stationInfo.RGT == 0 && stationInfo.TG == 0 && stationInfo.TPower == 0 {
		return Station{}, errors.New("invalid params")
	}
	DB := orm.NewOrm()
	p2 := Station{Name: stationInfo.Name}
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		id, err := DB.Insert(&stationInfo)
		if err != nil {
			fmt.Println("insert stationInfo error: ", err)
			return Station{}, err
		}
		stationInfo.Id = int(id)
		return stationInfo, nil
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return Station{}, err
	}
	if stationInfo.TPower != 0{
		p2.TPower = stationInfo.TPower
	}
	if stationInfo.TG != 0{
		p2.TG = stationInfo.TG
	}
	if stationInfo.RGT != 0 {
		p2.RGT = stationInfo.RGT
	}
	if stationInfo.Diameter != 0 {
		p2.Diameter = stationInfo.Diameter
	}
	DB.Update(&p2)
	return p2, nil
}
