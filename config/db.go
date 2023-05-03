package config

import (
	"fmt"
	"os"

	"github.com/usvishnuprakash/gin-grom-restful-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	var dns string

	dns = "user=postgres password=12345678 dbname=gotest port=5432 sslmode=disable"

	if os.Getenv("ENV") == "DEV" {
		dns = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	}
	fmt.Println("env", os.Getenv("ENV"))

	var db *gorm.DB

	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
		db.Exec("CREATE DATABASE gotest")

		fmt.Println(err)

	}

	if err != nil {
		fmt.Println("db connection error")
		panic(err)
	}

	fmt.Println("db connection established")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Campaign{})
	DB = db
}
