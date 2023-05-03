package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/usvishnuprakash/gin-grom-restful-api/config"
	"github.com/usvishnuprakash/gin-grom-restful-api/routes"
)

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env.example")
	err = godotenv.Load(environmentPath)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()

	// resting api

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routes.UserRoute(r)
	routes.CampaignRoutes(r)
	config.Connect()
	r.Run(":3000")

}
