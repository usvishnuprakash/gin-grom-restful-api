package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usvishnuprakash/gin-grom-restful-api/config"
	"github.com/usvishnuprakash/gin-grom-restful-api/models"
)

func GetCampaignList(c *gin.Context) {
	campaignList := []models.Campaign{}
	config.DB.Find(&campaignList)
	c.JSON(http.StatusOK, &campaignList)
}

func CreateCampaigns(c *gin.Context) {
	var campaignsList []models.Campaign
	c.BindJSON(&campaignsList)
	config.DB.Create(&campaignsList)
	c.JSON(http.StatusOK, &campaignsList)
}
