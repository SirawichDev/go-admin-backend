package main

import (
	"admin/src/database"
	"admin/src/models"

	"github.com/bxcodec/faker/v4"
	"github.com/segmentio/ksuid"
)

func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		secureId := ksuid.New()
		ambassadors := models.User{
			Id:           secureId.String(),
			Firstname:    faker.FirstName(),
			Lastname:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}
		ambassadors.SetPassword("1w2e3r4t5y")
		database.DB.Create(&ambassadors)
	}
}
