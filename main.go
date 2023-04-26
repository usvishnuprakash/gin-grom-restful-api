package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usvishnuprakash/gin-grom-restful-api/config"
	"github.com/usvishnuprakash/gin-grom-restful-api/routes"
)

func main() {

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
