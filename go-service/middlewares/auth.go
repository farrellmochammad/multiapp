package middlewares

import (
	"fmt"
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

		// set User Id Variable
		c.Set("user_id", valid.Claims.(jwt.MapClaims)["user_id"])
		fmt.Println(valid.Claims.(jwt.MapClaims)["user_id"], valid.Claims.(jwt.MapClaims)["exp"])
		c.Next()
	}
}
