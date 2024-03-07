package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/SamPariat/expenses-tracker/internal/pkg/dotenv"
	"github.com/SamPariat/expenses-tracker/internal/register"
)

func main() {
	router := gin.Default()

	err := dotenv.LoadEnvironment()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	allControllers := register.InitApp()

	serverRun := router.Group("/")
	{
		serverRun.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Server is running"})
		})
	}

	user := router.Group("/users")
	{
		user.GET("", allControllers.UserController.GetUsers)
		user.GET("/:id", allControllers.UserController.GetUser)
		user.POST("", allControllers.UserController.CreateUser)
		user.PUT("/:id", allControllers.UserController.UpdateUser)
		user.DELETE("/:id", allControllers.UserController.DeleteUser)
	}

	err = router.Run(":9623")
	if err != nil {
		log.Fatal("Error running server", err)
	}
}
