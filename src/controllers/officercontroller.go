package controllers

import (
	"net/http"
	"server/lib"
	"server/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

type OfficialController struct{}

func (OfficialController) Get(ctx *gin.Context) {
	now := time.Now()

	var officials []models.Official
	if err := lib.Database.Find(&officials).Where("start_term > ? AND ? < end_term", now, now).Preload("Resident").Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"officials": officials})
}
