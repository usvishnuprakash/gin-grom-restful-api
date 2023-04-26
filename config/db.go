package config

import (
	"fmt"

	"github.com/usvishnuprakash/gin-grom-restful-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dns := "user=postgres password=12345678 dbname=gotest port=5432 sslmode=disable"
	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		db, err = gorm.Open(postgres.Open("user=postgres password=12345678 dbname=gotest port=5432 sslmode=disable host=/private/tmp"), &gorm.Config{})
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
