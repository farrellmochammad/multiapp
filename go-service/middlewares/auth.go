package middlewares

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		// Check if toke in correct format
		// ie Bearer: xx03xllasx
		b := "Bearer: "
		if !strings.Contains(token, b) {
			c.JSON(401, gin.H{"message": "Gagal login, token tidak valid"})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(401, gin.H{"message": "Authorisasi tidak valid"})
			c.Abort()
			return
		}
		// Validate token
		valid, err := ValidateToken(t[1], SigningKey)
		if err != nil {
			c.JSON(401, gin.H{"message": "Token expired, silahkan login kembali"})
			c.Abort()
			return
		}

		c.Set("phone", valid.Claims.(jwt.MapClaims)["phone"])
		c.Set("name", valid.Claims.(jwt.MapClaims)["name"])
		c.Set("role", valid.Claims.(jwt.MapClaims)["role"])

		c.Next()
	}
}
