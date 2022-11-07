package database

import (
	"admin/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("exy:exypassword@tcp(db:3306)/gomin"), &gorm.Config{})
	if err != nil {
		log.Panic("Could not connect to database")
	}
}

func AutoMigrate()  {
	DB.AutoMigrate(models.User{})
}