package controllers

import (
	"net/http"
	"server/lib"
	"server/src/models"

	"github.com/gin-gonic/gin"
)

type ResidentController struct{}

func (ResidentController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id != "" {
		var resident models.Resident
		if err := lib.Database.First(&resident, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Resident not found"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"resident": resident})
	} else {
		var residents []models.Resident
		if err := lib.Database.Find(&residents).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"residents": residents})
	}
}
