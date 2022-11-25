package main

import (
	"admin/src/database"
	"admin/src/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var err error
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	err = app.Listen(":8000")
	if err != nil {
		log.Panic("cannot serve server")
	}
}
