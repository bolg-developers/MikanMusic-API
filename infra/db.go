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
	db, err := gorm.Open("mysql",
		config.Env().DBUser+":"+config.Env().DBPassword+
			"@tcp("+config.Env().DBAddress+":"+config.Env().DBPort+")/"+
			config.Env().DBName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	_db = db
}
