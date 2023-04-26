package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/usvishnuprakash/gin-grom-restful-api/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.UserController)
	router.POST("/user", controller.AddUser)

}
