package routes

import (
	"github.com/gofiber/fiber/v2"
	"webapi/daily-ping-api/app/controllers"
)


func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get_user", controllers.GetUser)
	api.Post("/create_user", controllers.CreateUser)
}