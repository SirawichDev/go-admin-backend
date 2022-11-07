package controllers

import (
	"admin/src/database"
	"admin/src/models"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/segmentio/ksuid"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	secureId := ksuid.New()

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password not match",
		})
	}
	user := models.User{
		Id:        secureId.String(),
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
		IsAdmin:   false,
	}
	user.SetPassword(data["password"])
	database.DB.Save(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == "0" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password not correct",
		})
	}
	payload := jwt.StandardClaims{
		Subject:   user.Id,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credential",
		})
	}
	cookie := fiber.Cookie{
		Name: "token-x",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "login success",
	})
}
