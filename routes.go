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
			resident.POST("", controller.Resident.Post)
			resident.PATCH(":id", controller.Resident.Patch)
			resident.DELETE("", controller.Resident.Delete)
		}
		event := api.Group("/events")
		{
			event.GET("", controller.Event.Get)
			event.GET("/:id", controller.Event.Get)
			event.POST("", controller.Event.Post)
			event.PATCH("/:id", controller.Event.Patch)
			event.DELETE("", controller.Event.Delete)
		}
	}
}
