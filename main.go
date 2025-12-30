package main

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"webapi/daily-ping-api/utils"
	"webapi/daily-ping-api/models"
	"webapi/daily-ping-api/storage"
	"webapi/daily-ping-api/app/controllers"
)


func(r *utils.Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get_user", r.models.GetUser(r))
	api.Post("/create_user", r.models.CreateUser(r))
}

func main() {
	err := godotenv.Load(".env")
	if (err != nil) {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:	os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if(err != nil) {
		log.Fatal("There is an error with database")
	}

	err = models.MigrateUsers(db)

	if (err != nil) {
		log.Fatal("Could not migrate")
	}

	err = models.AlterUsersTable(db)

	if (err != nil) {
		log.Fatal("Could not run migration for alter users table")
	}

	
	r := Repository {
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}