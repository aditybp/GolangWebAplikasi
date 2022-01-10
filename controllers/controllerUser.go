package controllers

import (
	"AplikasiWeb/database"
	"AplikasiWeb/models"

	"github.com/gofiber/fiber/v2"
)

func ListUser(c fiber.Ctx) error {
	var datauser []models.User
	database.DB.Find(&datauser)

	if len(datauser) == 0 {
		return c.Status(404).JSON(fiber.Map{"Status": "eror", "data": nil})
	}

	return c.JSON(&datauser)
}
