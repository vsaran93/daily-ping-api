package main

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

type Repository struct {
	DB *gorm.DB
}

type HealthStatus struct {
	Status      string   `json:"status"`
	UserId      string   `json:"userid"`
}

func(r *Repository) UpdateStatus(context *fiber.Ctx) error {
	healthStatus := HealthStatus{}

	err := context.BodyParser(&healthStatus)

	if(err != nil) {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request Failed"})
			return err
	}

	err := r.DB.Create(healthStatus).Error 

	if (err != nil) {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Bad Request"})
			return err
	}
	context.Status(http.StatusOk).JSON(&fiber.Map{"message": "Status has been added"})
	return nil

}

func(r *Repository) SetupRoutes(app *fiber.app) {
	api := app.group("/api")
	api.post("/update_status", r.UpdateStatus)
}

func main() {
	err := godotenv.Load(".env")
	if (err != nil) {
		log.Fatal(err)
	}

	r := Repository {
		DB: db,
	}

	db, err := storage.NewConnection(config)

	if(err != nil) {
		log.Fatal("There is an error with database")
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}