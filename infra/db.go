package infra

import (
	"github.com/bolg-developers/MikanMusic-API/config"
	"github.com/jinzhu/gorm"

	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

var _db *gorm.DB

func DB() *gorm.DB {
	return _db
}

func init() {
	db, err := gorm.Open(config.Env().DB[:5], config.Env().DB[8:])
	if err != nil {
		panic(err)
	}
	_db = db
}
