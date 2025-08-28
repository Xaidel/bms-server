package main

import (
	"fmt"
	"server/config"
	"server/lib"
	"server/src/models"
)

func init() {
	config.Load()
	lib.ConnectDatabase()
}

func main() {
	if err := lib.Database.AutoMigrate(
		&models.Resident{},
		&models.Household{},
		&models.ResidentHousehold{},
		&models.Health{},
		&models.Official{},
		&models.Certificate{},
		&models.Blotter{},
		&models.Event{},
		&models.Expense{},
		&models.Income{},
		&models.Setting{},
		&models.User{},
		&models.Mapping{},
	); err != nil {
		fmt.Println("Error Migrating")
		return
	}

	fmt.Println("Tables successfully migrated")
}
