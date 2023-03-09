package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DBConnect() {
	var err error
	dsn := os.Getenv("DB_DNS")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect to database")
	} else {
		fmt.Println("Connected to ozamotDB")
	}
}
