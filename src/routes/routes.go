package routes

import (
	"admin/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	admin.Get("/profile", controllers.User)
	admin.Post("/logout", controllers.Logout)
}
