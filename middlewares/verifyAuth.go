package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"sharely/configs"
	"sharely/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyAuth(c *gin.Context) {
	// get the token from cookie
	// accessToken, err := c.Cookie("Authorization")
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// }
	if c.Request.Header["Token"] != nil {
		// decoded token
		token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_TOKEN")), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// check the exp
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// find user with the token
			var user models.User
			configs.DB.First(&user, claims["sub"])
			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// attach to req
			c.Set("user", user)

			// continue
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
