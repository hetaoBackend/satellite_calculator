package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
	p2 := Station{Name: name}
	// 查询
	DB := orm.NewOrm()
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		logs.Error("查询不到")
		return &Station{}, err
	} else if err == orm.ErrMissPK {
		logs.Error("找不到主键")
		return &Station{}, err
	} else {
		logs.Info(p2)
	}
	return &p2, nil
}

func AddStation(stationInfo Station) (Station, error) {
	if stationInfo.Name == "" {
		return Station{}, errors.New("invalid params")
	}
	if stationInfo.Diameter == 0 && stationInfo.RGT == 0 && stationInfo.TG == 0 && stationInfo.TPower == 0 {
		return Station{}, errors.New("invalid params")
	}
	DB := orm.NewOrm()
	p2 := Station{Name: stationInfo.Name}
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		id, err := DB.Insert(&stationInfo)
		if err != nil {
			logs.Info("update station info, name: ", stationInfo.Name)
			logs.Error("insert stationInfo error: ", err)
			return Station{}, err
		}
		stationInfo.Id = int(id)
		return stationInfo, nil
	} else if err == orm.ErrMissPK {
		logs.Error("找不到主键")
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
	logs.Info("update station info, name: ", stationInfo.Name)
	DB.Update(&p2)
	return p2, nil
}
