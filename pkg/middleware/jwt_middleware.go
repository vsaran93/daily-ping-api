package middleware 

import (
	"github.com/gofiber/fiber/v2"
	"os"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)


func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey: "jwt",
		ErrorHandler: jwtError,
	}
	return jwtMiddleware.New(config)
}	


func jwtError(context *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}
	return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": true,
		"msg": err.Error(),
	})
}