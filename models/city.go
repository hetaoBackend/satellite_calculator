package models

import "fmt"

type CityInfo struct {
	Name string
	Longtitude float32
	Latitude float32
}

// ToDo: add city info to DB
func AddCity(cityInfo CityInfo) (int, error) {
	fmt.Println(cityInfo)
	return 1, nil
}
