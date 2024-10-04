package user

import (
	"MyHelp/db"
	"MyHelp/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		// 未登录情况
		if err != nil {
			c.Set("role", models.Guest)
			c.Next()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(models.Conf.JwtSecret), nil
		})

		if err != nil {
			//c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't parse token"})
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var user models.User
			user.Id = int(claims["id"].(float64))
			user.Role = models.Role(claims["role"].(float64))
			c.Set("id", user.Id)
			c.Set("role", user.Role)
			err := db.SelectUserProfileById(&user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
				return
			}
			c.Set("username", user.Username)
			c.Next()
		} else {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
	}
}
