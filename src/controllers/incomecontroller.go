package controllers

import (
	"net/http"
	"server/lib"
	"server/src/models"

	"github.com/gin-gonic/gin"
)

type IncomeController struct{}

func (IncomeController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	if id != "" {
		var income models.Income
		if err := lib.Database.First(&income, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Income not found"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"incomes": income})
	} else {
		var incomes []models.Income
		if err := lib.Database.Find(&incomes).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"incomes": incomes})
	}
}

func (IncomeController) Post(ctx *gin.Context) {
}
