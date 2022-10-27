package routers

import (
	"final-project-go/handlers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", handlers.UserRegister)
		userRouter.POST("/login", handlers.UserLogin)
	}

	return router
}