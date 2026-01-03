package routes

import (
	"github.com/gofiber/fiber/v2"
	"webapi/daily-ping-api/app/controllers"
)


func PublicRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/token/new", controllers.GenerateNewAccessToken)
	api.Post("/create_user", controllers.CreateUser)
}