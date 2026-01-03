package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"webapi/daily-ping-api/models"
	"webapi/daily-ping-api/storage"
	"webapi/daily-ping-api/pkg/routes"
	"github.com/joho/godotenv"
)




func main() {
	err := godotenv.Load(".env")
	if (err != nil) {
		log.Fatal(err)
	}
	
	db := storage.OpenDbConnection()

	err = models.MigrateUsers(db)

	if (err != nil) {
		log.Fatal("Could not migrate")
	}

	err = models.AlterUsersTable(db)

	if (err != nil) {
		log.Fatal("Could not run migration for alter users table")
	}

	sqlDB, err := db.DB()
	sqlDB.Close();

	app := fiber.New()
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	app.Listen(":8080")
}