package main

import (
	"server/src/controllers"

	"github.com/gin-gonic/gin"
)

func APIRoutes(router *gin.Engine) {
	controller := &controllers.Controller{}

	api := router.Group("api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("login", controller.Auth.Login)
		}
		resident := api.Group("/residents")
		{
			resident.GET("", controller.Resident.Get)
			resident.GET("/:id", controller.Resident.Get)
		}
	}
}
