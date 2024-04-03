package main

import (
	"crud-operations/config"
	"crud-operations/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8000")
}
