package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Satellite struct {
	Id        int `orm:"pk"`
	Name      string
	FreType   string
	GT        float32
	SFD       float32
	EIRP      float32
	Longitude float32
}

func GetSatellite(name, freType string) (*Satellite, error) {
	if name == "" || freType == "" {
		return &Satellite{}, errors.New("invalid params")
	}
	p2 := Satellite{Name: name, FreType:freType}
	// 查询
	DB := orm.NewOrm()
	err := DB.Read(&p2, "Name", "FreType")
	if err == orm.ErrNoRows {
		logs.Error("查询不到")
		return &Satellite{}, err
	} else if err == orm.ErrMissPK {
		logs.Error("找不到主键")
		return &Satellite{}, err
	} else {
		logs.Info(p2)
	}
	return &p2, nil
}

func AddSate(sateInfo Satellite) (Satellite, error) {
	if sateInfo.Name == "" || sateInfo.FreType == "" {
		return Satellite{}, errors.New("invalid params")
	}
	if sateInfo.EIRP == 0 && sateInfo.Longitude == 0 && sateInfo.GT == 0 && sateInfo.SFD == 0 {
		return Satellite{}, errors.New("invalid params")
	}
	DB := orm.NewOrm()
	p2 := Satellite{Name: sateInfo.Name, FreType: sateInfo.FreType}
	err := DB.Read(&p2, "Name", "FreType")
	if err == orm.ErrNoRows {
		logs.Info("insert satellite info, name: ", sateInfo.Name)
		id, err := DB.Insert(&sateInfo)
		if err != nil {
			logs.Error("insert sateInfo error: ", err)
			return Satellite{}, err
		}
		sateInfo.Id = int(id)
		return sateInfo, nil
	} else if err == orm.ErrMissPK {
		logs.Error("找不到主键")
		return Satellite{}, err
	}
	if sateInfo.GT != 0 {
		p2.GT = sateInfo.GT
	}
	if sateInfo.SFD != 0{
		p2.SFD = sateInfo.SFD
	}
	if sateInfo.EIRP != 0{
		p2.EIRP = sateInfo.EIRP
	}
	if sateInfo.Longitude != 0{
		p2.Longitude = sateInfo.Longitude
	}
	logs.Info("update satellite info, name: ", sateInfo.Name)
	DB.Update(&p2)
	return p2, nil
}
