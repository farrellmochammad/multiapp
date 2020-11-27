package usecase

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"go-service/middlewares"
	models "go-service/model"
	repository "go-service/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func DeliveryAuth(r *gin.RouterGroup) {

	r.POST("/register", func(c *gin.Context) {
		var userReqister models.User
		err := c.Bind(&userReqister)

		if err != nil {
			c.JSON(500, gin.H{
				"status": "failed",
				"Error":  "There is error on query string",
			})
			return
		}

		rand.Seed(time.Now().UnixNano())

		chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"abcdefghijklmnopqrstuvwxyz" +
			"0123456789")
		length := 4
		var b strings.Builder
		for i := 0; i < length; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		password := b.String()

		fmt.Println(password)

		user := userReqister
		user.Password = password

		stat := repository.InsertUser(user)

		if stat {
			c.JSON(200, gin.H{
				"status":   "success",
				"password": password,
			})
		} else {
			c.JSON(403, gin.H{
				"status": "failed",
			})
		}

	})

	r.POST("/userinfo", func(c *gin.Context) {
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
		valid, err := middlewares.ValidateToken(t[1], middlewares.SigningKey)
		if err != nil {
			c.JSON(401, gin.H{"message": "Token expired, silahkan login kembali"})
			c.Abort()
			return
		}

		// set User Id Variable
		c.JSON(200, gin.H{
			"name":  valid.Claims.(jwt.MapClaims)["Phone"],
			"phone": valid.Claims.(jwt.MapClaims)["Name"],
			"role":  valid.Claims.(jwt.MapClaims)["Role"],
		})

	})

	r.POST("/login", func(c *gin.Context) {
		var userLogin models.User
		c.Bind(&userLogin)

		// if &userLogin.Phone != nil {
		// 	c.JSON(500, gin.H{
		// 		"status": "failed",
		// 		"Error":  "There is error on query string",
		// 	})
		// 	return
		// }

		user := repository.CheckUserLogin(userLogin).(models.User)
		switch user.Name {
		case "-1":
			c.JSON(403, gin.H{
				"status": "Not success",
				"token":  "Username or password not valid",
			})
		case "-2":
			c.JSON(500, gin.H{
				"status": "Not success",
				"token":  "Username and password not match",
			})
		default:
			token, _ := middlewares.GenerateToken([]byte(middlewares.SigningKey), user)
			c.JSON(200, gin.H{
				"status": "login success",
				"token":  token,
			})
		}

	})

}
