package db

import (
	"time"

	_ "github.com/lib/pq" // postgres driver
	"xorm.io/core"
	"xorm.io/xorm"
)

// NewDBEngine --
var engine *xorm.Engine

// InitDBEngine -
func InitDBEngine(driver, source string) (err error) {
	engine, err = xorm.NewEngine(driver, source)
	if err != nil {
		panic("Failed to connect database")
	}
	engine.SetTZDatabase(time.Local)
	engine.SetTZLocation(time.Local)
	err = engine.Ping()
	if err != nil {
		panic("Failed to connect database")
	}
	engine.ShowSQL(false)
	engine.SetColumnMapper(core.GonicMapper{})

	err = engine.Sync2(Fix{})
	return err

}
