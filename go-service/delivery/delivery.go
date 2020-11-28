package delivery

import (
	"go-service/middlewares"
	usecase "go-service/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	v1 := r.Group("api/v1")

	v1.Use(cors.Default())
	{
		usecase.DeliveryAuth(v1)
	}
	v1.Use(middlewares.AuthHandler())
	{
		usecase.DeliveryArea(v1)
	}

	return r

}
