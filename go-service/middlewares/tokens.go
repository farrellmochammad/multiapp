package middlewares

import (
	models "go-service/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(k []byte, user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["Phone"] = user.Phone
	claims["Name"] = user.Name
	claims["Role"] = user.Role
	claims["Password"] = user.Password
	claims["Timestamp"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(k)
	return tokenString, err
}

func ValidateToken(t string, k string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(k), nil
	})

	return token, err
}
