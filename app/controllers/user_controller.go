package controllers


import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"webapi/daily-ping-api/models"
	"webapi/daily-ping-api/storage"
)

type User struct {
	Phone			string		`json:"phone"`
	IsPhoneVerified	bool  		`json:"is_phone_verified"`
}


func CreateUser(context *fiber.Ctx) error {
	user := User{}

	db := storage.OpenDbConnection();
	err := context.BodyParser(&user)

	if(err != nil) {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request Failed"})
			return err
	}

	err = db.Create(user).Error 

	if (err != nil) {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Bad Request"})
			return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "User has been added"})
	return nil

}

func GetUser(context *fiber.Ctx) error {
	users := []models.User{}

	db := storage.OpenDbConnection();

	db.Find(&users)
	if (len(users) == 0) {
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "Internal server error"})
			return nil
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"success": true, "data": users, })
		return nil
}