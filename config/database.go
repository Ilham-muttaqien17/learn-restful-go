package config

import (
	"fmt"

	"github.com/Ilham-muttaqien17/learn-restful-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Env.DBUsername, Env.DBPassword, Env.DBHost, Env.DBPort, Env.DBName)
	db, err := gorm.Open(mysql.Open(dsn))

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