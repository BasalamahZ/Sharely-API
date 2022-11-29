package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// sign and encoded token as a string
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		return accessToken, err
	}
	return accessToken, nil
}
