package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"math"
)

type CalculateRequest struct {
	SateName      string
	UpFrequency   float32
	DownFrequency float32
	FreType       string
	TStationName  string
	RStationName  string
	TCityName     string
	RCityName     string
	BOo           float32
	BOi           float32
	DecThreshold  float32
}

type CalculateResponse struct {
	UpCN0       float64
	DownCN0     float64
	TotalCN0    float64
	PowerRatio  float64
	MaxBitRate  float64
}

func Calculate(calculateInfo CalculateRequest) (CalculateResponse, error) {
	if calculateInfo.SateName == "" || calculateInfo.TStationName == "" || calculateInfo.RStationName == "" {
		return CalculateResponse{}, errors.New("invalid params")
	}
	sateInfo, err := GetSatellite(calculateInfo.SateName, calculateInfo.FreType)
	if err != nil {
		logs.Error("Get satellite info err: ", err)
		return CalculateResponse{}, err
	}
	tStationInfo, err := GetStation(calculateInfo.TStationName, calculateInfo.FreType)
	if err != nil {
		logs.Error("Get transmit station info err: ", err)
		return CalculateResponse{}, err
	}
	rStationInfo, err := GetStation(calculateInfo.RStationName, calculateInfo.FreType)
	if err != nil {
		logs.Error("Get receive station info err: ", err)
		return CalculateResponse{}, err
	}
	tCityInfo, err := GetCity(calculateInfo.TCityName)
	if err != nil {
		logs.Error("Get transmit station location info err: ", err)
		return CalculateResponse{}, err
	}
	rCityInfo, err := GetCity(calculateInfo.RCityName)
	if err != nil {
		logs.Error("Get receive station location info err: ", err)
		return CalculateResponse{}, err
	}

	logs.Info("[%s] The uplink frequency is %v, the downlink frequency is %v", calculateInfo.FreType, calculateInfo.UpFrequency, calculateInfo.DownFrequency)

	r := 42164.2
	Re := 6378.155
	tStationDis := math.Sqrt(r*r + Re * Re - 2*Re*r*math.Cos(float64(tCityInfo.Latitude/180) * math.Pi) * math.Cos(float64((tCityInfo.Longitude - sateInfo.Longitude)/180) * math.Pi))
	rStationDis := math.Sqrt(r*r + Re * Re - 2*Re*r*math.Cos(float64(rCityInfo.Latitude/180) * math.Pi) * math.Cos(float64((rCityInfo.Longitude - sateInfo.Longitude)/180) * math.Pi))

	upLoss := 92.45 + 20*math.Log10(float64(calculateInfo.UpFrequency)*tStationDis)
	downLoss := 92.45 + 20*math.Log10(float64(calculateInfo.DownFrequency)*rStationDis)

	tEIRP := tStationInfo.TPower + tStationInfo.TG

	upCN0 := float64(tEIRP + sateInfo.GT) - upLoss + 228.6
	flux := float64(tEIRP) - 10*math.Log10(2*math.Pi*tStationDis*tStationDis) - 60

	EIRPs := float64(sateInfo.EIRP - calculateInfo.BOo)

	powerRatio := float64(1)
	if ratio := flux - float64(sateInfo.SFD - calculateInfo.BOi); ratio < 0 {
		EIRPs += float64(ratio)
		powerRatio = math.Pow(10, ratio/10)
	}

	downCN0 :=  EIRPs + float64(rStationInfo.RGT) - downLoss + 228.6
	totalCN0 := -10 * math.Log10(math.Pow(10,-upCN0/10) + math.Pow(10, -downCN0/10))
	maxBitRate := math.Pow(10, (totalCN0 - float64(calculateInfo.DecThreshold))/10)
	logs.Info("uplink C/N0 is %v, downLink C/N0 is %v, total C/N0 is %v, powerRatio is %v, maxBitRate is %v", upCN0, downCN0, totalCN0, powerRatio, maxBitRate)
	return CalculateResponse{UpCN0:upCN0, DownCN0:downCN0, TotalCN0:totalCN0, PowerRatio:powerRatio, MaxBitRate:maxBitRate}, nil
}
