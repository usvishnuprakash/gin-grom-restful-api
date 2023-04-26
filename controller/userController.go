package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usvishnuprakash/gin-grom-restful-api/config"
	"github.com/usvishnuprakash/gin-grom-restful-api/models"
)

func UserController(c *gin.Context) {
	usersList := []models.User{}
	config.DB.Find(&usersList)

	c.JSON(http.StatusOK, &usersList)
}

func AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(http.StatusOK, &user)

}
