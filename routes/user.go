package routes

import (
	"crud-operations/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/", controller.DeleteUser)
	router.PUT("/", controller.UpdateUser)
}
