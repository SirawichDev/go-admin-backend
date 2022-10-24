package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	_, err := gorm.Open(mysql.Open("exy:exypassword@tcp(db:3306)/gomin"), &gorm.Config{})
	if err != nil {
		print("Could not connect to database")
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("fff")
	})
	err = app.Listen(":8000")
	if err != nil {
		log.Panic("cannot serve server")
	}
}
