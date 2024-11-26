package db

import "log"

type HazardPoint struct {
	ID        int64   `json:"id" xorm:"id"`
	Longitude float64 `json:"longitude" xorm:"longitude"`
	Latitude  float64 `json:"latitude" xorm:"latitude"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	County    string  `json:"county"`
	Town      string  `json:"town"`
	Group     string  `json:"group"`
	Location  string  `json:"location" xorm:"location"`
}

func List() []*HazardPoint {
	var x []*HazardPoint
	err := engine.Table(HazardPoint{}).Where("id = 2436").Limit(4000).OrderBy("id").Find(&x)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

type Fix struct {
	ID         int64  `json:"id" xorm:"id  pk autoincr"`
	HPID       int64  `json:"hp_id" xorm:"hp_id"`
	Province   string `json:"province"`
	City       string `json:"city"`
	County     string `json:"county"`
	Town       string `json:"town"`
	HPProvince string `json:"hp_province" xorm:"hp_province"`
	HPCity     string `json:"hp_city" xorm:"hp_city"`
	HPCounty   string `json:"hp_county" xorm:"hp_county"`
	HPTown     string `json:"hp_town" xorm:"hp_town"`
}

func InsertFix(x Fix) error {
	_, err := engine.Table(Fix{}).Insert(x)
	return err
}
