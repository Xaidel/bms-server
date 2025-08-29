package controllers

import (
	"net/http"
	"server/lib"
	"server/src/models"

	"github.com/gin-gonic/gin"
)

type MappingController struct{}

func (MappingController) Get(ctx *gin.Context) {
	var mappings []models.Mapping
	if err := lib.Database.Preload("Household.Residents.Resident").Find(&mappings).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mappings": mappings})
}
