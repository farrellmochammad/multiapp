package middlewares

import (
	models "go-service/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(k []byte, user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["phone"] = user.Phone
	claims["name"] = user.Name
	claims["role"] = user.Role
	claims["password"] = user.Password
	claims["timestamp"] = time.Now().Unix()
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
