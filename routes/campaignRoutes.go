package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/usvishnuprakash/gin-grom-restful-api/controller"
)

func CampaignRoutes(router *gin.Engine) {

	router.GET("/campaigns", controller.GetCampaignList)
	router.POST("/campaigns", controller.CreateCampaigns)
}
