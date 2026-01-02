package controllers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"webapi/daily-ping-api/utils"
)

func GenerateNewAccessToken(context *fiber.Ctx) error {
	token, err := utils.GenerateToken()

	if (err != nil) {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Could not generate token"})
			return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"success": true, "access_token": token, })
		return nil
}