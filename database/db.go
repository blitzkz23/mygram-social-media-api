package database

import (
	"fmt"
	"log"
	"mygram-social-media-api/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	username = "postgres"
	password = "naufalaldy23"
	host     = "localhost"
	dbPort   = "5432"
	dbName   = "mygram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, dbPort, username, dbName, password)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to database")
	db.Debug().AutoMigrate(entity.User{}, entity.Photo{}, entity.SocialMedia{}, entity.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
