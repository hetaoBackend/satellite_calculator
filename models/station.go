package models

import "fmt"

type StationInfo struct {
	Diameter float32
	TPower float32
	TG float32
	RGT float32
}

// ToDo: add station info to DB
func AddStation(sateInfo SatelliteInfo) (int, error) {
	fmt.Println(sateInfo)
	return 1, nil
}
