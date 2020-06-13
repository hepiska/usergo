package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hepiska/todo-go/controllers"
)

func globalroute(router *gin.Engine) {
	router.GET("/alive", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"meesage": "alive"})
	})
}

func authRoute(router *gin.Engine) {
	AuthController := new(controllers.AuthController)
	router.POST("/signup", AuthController.Signup)
}

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	authRoute(router)
	globalroute(router)
	return router
}
