package db

import (
	"example.com/m/v2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {

	USER := "root"
	PASS := "1234"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "2022_Hackerton"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.StockInfo{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}
