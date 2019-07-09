package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Satellite struct {
	Id        int `orm:"pk"`
	Name      string
	GT        float32
	SFD       float32
	EIRP      float32
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

func AddSate(sateInfo Satellite) (Satellite, error) {
	if sateInfo.EIRP == 0 && sateInfo.Longitude == 0 && sateInfo.GT == 0 && sateInfo.SFD == 0 {
		return Satellite{}, errors.New("invalid params")
	}
	DB := orm.NewOrm()
	p2 := Satellite{Name: sateInfo.Name}
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		id, err := DB.Insert(&sateInfo)
		if err != nil {
			fmt.Println("insert sateInfo error: ", err)
			return Satellite{}, err
		}
		sateInfo.Id = int(id)
		return sateInfo, nil
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
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
	DB.Update(&p2)
	return p2, nil
}
