package config

import (
	"github.com/Ilham-muttaqien17/learn-restful-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB() error {
	db, err := gorm.Open(mysql.Open("test:password@tcp(localhost:3306)/go_restapi"))

	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Book{})

	DB = db

	return nil
}

func DisconnectDB() error {
	conn, err := DB.DB()

	if err != nil {
		return err
	}

	conn.Close()
	return nil
}