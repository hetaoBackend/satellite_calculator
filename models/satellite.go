package models

import "fmt"

type SatelliteInfo struct {
	Name string
	GT float32
	SFD float32
	EIRP float32
	Longitude float32
}

// ToDo: add satellite info to DB
func AddSate(sateInfo SatelliteInfo) (int, error) {
	fmt.Println(sateInfo)
	return 1, nil
}