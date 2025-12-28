package main

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"webapi/daily-ping-api/models"
	"webapi/daily-ping-api/storage"
)

type Repository struct {
	DB *gorm.DB
}

type User struct {
	Phone			string		`json:"phone"`
	IsPhoneVerified	bool  		`json:"is_phone_verified"`
}

func(r *Repository) CreateUser(context *fiber.Ctx) error {
	user := User{}

	err := context.BodyParser(&user)

	if(err != nil) {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request Failed"})
			return err
	}

	err = r.DB.Create(user).Error 

	if (err != nil) {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Bad Request"})
			return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "User has been added"})
	return nil

}

func(r *Repository) GetUser(context *fiber.Ctx) error {
	users := []models.Users{}

	r.DB.Find(&users)
	if (len(users) == 0) {
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "Internal server error"})
			return nil
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"success": true, "data": users, })
		return nil
}

func(r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get_user", r.GetUser)
	api.Post("/create_user", r.CreateUser)
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