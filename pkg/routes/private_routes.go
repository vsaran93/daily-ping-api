package routes

import (
	"github.com/gofiber/fiber/v2"
	"webapi/daily-ping-api/app/controllers"
	"webapi/daily-ping-api/pkg/middleware"
)

func PrivateRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/user", middleware.JWTProtected(), controllers.GetUser)
}