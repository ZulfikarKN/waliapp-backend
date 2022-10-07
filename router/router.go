package router

import (
	"Go/WALIAPP-BACKEND/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	guru := router.Group("teacher")

	guru.POST("/registerAdmin", controller.SaveAdmin)
	guru.POST("/add", controller.SaveTeacher)

	return router
}