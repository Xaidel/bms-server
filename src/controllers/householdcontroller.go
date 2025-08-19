package controllers

import (
	"fmt"
	"net/http"
	"server/lib"
	"server/src/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HouseholdController struct{}

type MemberDTO struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

type CreateHouseholdDTO struct {
	DateOfResidency time.Time   `json:"dateOfResidency"`
	HouseholdNumber string      `json:"householdNumber"`
	HouseholdType   string      `json:"householdType"`
	Members         []MemberDTO `json:"members"`
	Status          string      `json:"status"`
	Zone            string      `json:"zone"`
}

func (HouseholdController) Post(ctx *gin.Context) {
	var input CreateHouseholdDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := lib.Database

	err := db.Transaction(func(tx *gorm.DB) error {
		household := models.Household{
			HouseholdNumber: input.HouseholdNumber,
			Zone:            input.Zone,
			Type:            input.HouseholdType,
			Status:          input.Status,
			DateOfResidency: input.DateOfResidency,
		}

		if err := tx.Create(&household).Error; err != nil {
			return err
		}

		for _, m := range input.Members {
			// Check if resident already linked
			var existing models.ResidentHousehold
			if err := tx.Where("resident_id = ?", m.ID).First(&existing).Error; err == nil {
				// Load resident info
				var resident models.Resident
				if err := tx.First(&resident, m.ID).Error; err != nil {
					return fmt.Errorf("resident with ID %d already belongs to a household", m.ID)
				}

				// Construct full name with optional middlename and suffix
				fullName := resident.Firstname
				if resident.Middlename != nil && *resident.Middlename != "" {
					fullName += " " + *resident.Middlename
				}
				fullName += " " + resident.Lastname
				if resident.Suffix != nil && *resident.Suffix != "" {
					fullName += " " + *resident.Suffix
				}

				return fmt.Errorf("resident %s already belongs to a household", fullName)
			} else if err != gorm.ErrRecordNotFound {
				return err
			}

			rh := models.ResidentHousehold{
				HouseholdID: household.ID,
				ResidentID:  m.ID,
				Role:        m.Role,
			}
			if err := tx.Create(&rh).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "household created successfully",
	})
}
