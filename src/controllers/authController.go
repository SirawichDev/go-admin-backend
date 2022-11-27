package controllers

import (
	"admin/src/database"
	"admin/src/middlewares"
	"admin/src/models"
	"time"

	"github.com/dgrijalva/jwt-go"

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
		Id:           secureId.String(),
		Firstname:    data["first_name"],
		Lastname:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: false,
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
		Name:     "token-x",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "login success",
	})
}

func User(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token-x",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "logout successfully",
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	id, _ := middlewares.GetUserId(c)

	user = models.User{
		Id:        id,
		Email:     data["email"],
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
	}

	database.DB.Where("id = ?", id).Updates(&user)
	return c.JSON(&user)
}

func SetUpNewPassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password not match",
		})
	}
	id, _ := middlewares.GetUserId(c)
	var user models.User
	user = models.User{
		Id: id,
	}
	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(&user)
	return c.JSON(
		fiber.Map{
			"message": "new password config successfully!",
		},
	)

}
