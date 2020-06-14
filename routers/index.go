package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hepiska/todo-go/controllers"
	"github.com/hepiska/todo-go/middlewares"
	"github.com/hepiska/todo-go/utils"
)

func globalroute(router *gin.Engine) {
	router.GET("/alive", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"meesage": "alive"})
	})
}

func authRoute(router *gin.Engine) {
	AuthController := new(controllers.AuthController)
	router.POST("/signup", AuthController.Signup)
	router.POST("/login", AuthController.Login)
	authGroup := router.Group("/")
	authGroup.Use(middlewares.Authentication())
	authGroup.GET("/me", AuthController.GetMe)
}

func userRoute(router *gin.Engine) {
	userController := new(controllers.UserController)

	authGroup := router.Group("/users")
	authGroup.Use(middlewares.Authentication())
	authGroup.GET("/", userController.GetAll)
	authGroup.GET("/:id", userController.GetOne)
	authGroup.PUT("/:id", userController.Update)

	authGroup.DELETE("/:id", userController.Delete)

}

// InitRoute ins function to initial http route
func InitRoute() *gin.Engine {
	mode := utils.EnvVar("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)

	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authRoute(router)
	userRoute(router)
	globalroute(router)
	return router
}
