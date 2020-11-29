package usecase

import (
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
			c.JSON(401, gin.H{
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

		user := userReqister
		user.Password = password

		stat := repository.InsertUser(user)

		if stat {
			c.JSON(201, gin.H{
				"status":   "success",
				"password": password,
			})
		} else {
			c.JSON(401, gin.H{
				"status":  "failed",
				"message": "handphone number already exist",
			})
		}

	})

	r.GET("/userinfo", func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		b := "Bearer: "
		if !strings.Contains(token, b) {
			c.JSON(401, gin.H{"status": "failed", "message": "login failed, token not valid"})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(401, gin.H{"status": "failed", "message": "Authroization not valid"})
			c.Abort()
			return
		}
		// Validate token
		valid, err := middlewares.ValidateToken(t[1], middlewares.SigningKey)
		if err != nil {
			c.JSON(401, gin.H{"status": "failed", "message": "Token expired, please login back"})
			c.Abort()
			return
		}

		c.JSON(200, gin.H{
			"name":  valid.Claims.(jwt.MapClaims)["name"],
			"phone": valid.Claims.(jwt.MapClaims)["phone"],
			"role":  valid.Claims.(jwt.MapClaims)["role"],
		})

	})

	r.POST("/login", func(c *gin.Context) {
		var userLogin models.User
		c.Bind(&userLogin)

		user := repository.CheckUserLogin(userLogin).(models.User)
		switch user.Name {
		case "-1":
			c.JSON(401, gin.H{
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
