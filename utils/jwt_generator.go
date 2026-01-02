
package utils

import (
	"os"
	"github.com/golang-jwt/jwt"
	"time"
	"strconv"
)

func GenerateToken() (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRY_MINUTES"))

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if (err != nil) {
		return "", err
	}

	return t, nil
}