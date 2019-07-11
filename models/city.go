package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id        int `orm:"pk"`
	Name      string
	Longitude float32
	Latitude  float32
}

func GetCity(name string) (*City, error) {
	if name == "" {
		return &City{}, errors.New("invalid params")
	}
	p2 := City{Name: name}
	// 查询
	DB := orm.NewOrm()
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		logs.Error("查询不到")
		return &City{}, err
	} else if err == orm.ErrMissPK {
		logs.Error("找不到主键")
		return &City{}, err
	} else {
		logs.Info(p2)
	}
	return &p2, nil
}

func GetCities(offset, count int32) ([]*City, error) {
	if offset < 0 || count <= 0 {
		return nil, errors.New("invalid params")
	}
	cities := make([]*City, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("city")
	num, err := qs.Limit(count, offset).All(&cities)
	if err != nil {
		logs.Error("Get cities error, error: ", err)
		return nil, err
	}
	logs.Info("Get %v city rows", num)
	return cities, nil
}

func DeleteCity(name string) error {
	if name == "" {
		return errors.New("invalid params")
	}
	p2 := City{Name: name}
	// 查询
	DB := orm.NewOrm()
	num, err := DB.Delete(&p2, "Name")
	if err == orm.ErrNoRows {
		logs.Error("查询不到")
		return err
	} else if err == orm.ErrMissPK {
		logs.Error("找不到主键")
		return err
	} else {
		logs.Info("delete %v city rows", num)
	}
	return nil
}

func AddCity(cityInfo City) (City, error) {
	if cityInfo.Name == "" {
		return City{}, errors.New("invalid params")
	}
	if cityInfo.Longitude == 0 && cityInfo.Latitude == 0 {
		return City{}, errors.New("invalid params")
	}
	DB := orm.NewOrm()
	p2 := City{Name: cityInfo.Name}
	err := DB.Read(&p2, "Name")
	if err == orm.ErrNoRows {
		logs.Info("insert city info, name: ", cityInfo.Name)
		id, err := DB.Insert(&cityInfo)
		if err != nil {
			logs.Error("insert cityInfo error: ", err)
			return City{}, err
		}
		cityInfo.Id = int(id)
		return cityInfo, nil
	} else if err == orm.ErrMissPK {
		logs.Error("找不到主键")
		return City{}, err
	}
	if cityInfo.Latitude != 0 {
		p2.Latitude = cityInfo.Latitude
	}
	if cityInfo.Longitude != 0 {
		p2.Longitude = cityInfo.Longitude
	}
	logs.Info("update city info, name: ", cityInfo.Name)
	DB.Update(&p2)
	return p2, nil
}
