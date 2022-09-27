package database

import (
	"fmt"
	"log"

	"citybiker-go-api/db/config"
	"citybiker-go-api/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Hello() string {
	message := "Hi, Welcome!"
	return message
}

func Init() *gorm.DB {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln(err)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Station{}, &models.Trip{})
	fmt.Println("Database connected")

	return db
}
