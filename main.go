package main

import (
	"net/http"
	"server/config"
	"server/lib"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Load()
	lib.ConnectDatabase()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.1.2:1420", "http://localhost:1420", "http://192.168.123.39:1420", "http://192.168.123.35:1420"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "JSSL"})
	})
	APIRoutes(router)
	router.Run()
}
