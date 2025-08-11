package main

import "server/src/models"

var Users = []models.User{
	{
		Role:     "Barangay Captain",
		Username: "Brgy. Captain",
		Password: "admin",
	},
	{
		Role:     "Barangay Secretary",
		Username: "Brgy. Secretary",
		Password: "secretary",
	},
	{
		Role:     "Barangay Treasurer",
		Username: "Brgy. Treasurer",
		Password: "treasurer",
	},
}
