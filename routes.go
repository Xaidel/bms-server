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
		household := api.Group("/households")
		{
			household.GET("", controller.Household.Get)
			household.GET("/:id", controller.Household.GetOne)
			household.POST("", controller.Household.Post)
			household.DELETE("", controller.Household.Delete)
		}
		income := api.Group("/incomes")
		{
			income.GET("", controller.Income.Get)
			income.GET("/:id", controller.Income.Get)
			income.POST("", controller.Income.Post)
			income.DELETE("", controller.Income.Delete)
		}
		expense := api.Group("/expenses")
		{
			expense.GET("", controller.Expense.Get)
			expense.GET("/:id", controller.Expense.Get)
			expense.POST("", controller.Expense.Post)
			expense.DELETE("", controller.Expense.Delete)
		}
		logbook := api.Group("/logbooks")
		{
			logbook.GET("", controller.Logbook.Get)
			logbook.GET("/:id", controller.Logbook.Get)
			logbook.POST("", controller.Logbook.Post)
			logbook.PATCH("/:id", controller.Logbook.Patch)
			logbook.DELETE("", controller.Logbook.Delete)
		}
		official := api.Group("/officials")
		{
			official.GET("", controller.Official.Get)
			official.GET("/:id", controller.Official.Get)
			official.POST("", controller.Official.Post)
			official.PATCH("/:id", controller.Official.Patch)
			official.DELETE("", controller.Official.Delete)
		}
		mapping := api.Group("/mappings")
		{
			mapping.GET("", controller.Mapping.Get)
		}
	}
}
