package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/time/rate"

	"dizai/db"
	"dizai/gaode"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	master := "host=127.0.0.1 port=41116 user=meihuan password=meihuan dbname=dizai sslmode=disable"
	err := db.InitDBEngine("postgres", master)

	if err != nil {
		log.Fatal(err)
	}
	err = gaode.Init("3d96eda8c20e02f8cc4ea69635e2a510")
	if err != nil {
		log.Fatal(err)
	}
	x := db.List()
	limit := rate.NewLimiter(rate.Limit(3), 1)
	for _, point := range x {
		_ = limit.Wait(context.Background())
		p, c, d, t, err := gaode.InChina(point.Latitude, point.Longitude)
		if err != nil {
			fmt.Printf("id :%d  err:%s\n", point.ID, err.Error())
			continue
		}
		if p == "" {
			fmt.Printf("id :%d  err:%s\n", point.ID, "返回为空")
			continue
		}
		err = db.InsertFix(db.Fix{
			HPID:       point.ID,
			HPProvince: point.Province,
			HPCity:     point.City,
			HPCounty:   point.County,
			HPTown:     point.Town,
			Province:   p,
			City:       c,
			County:     d,
			Town:       t,
		})
		if err != nil {
			fmt.Printf("id :%d  err:%s\n", point.ID, err.Error())
		}
	}
}

// 2436
