package controllers

import (
	"admin/src/database"
	"admin/src/models"

	"github.com/gofiber/fiber/v2"
)

func Ambassadors(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Where("is_ambassador").Find(&users)

	return c.JSON(&users)
}
