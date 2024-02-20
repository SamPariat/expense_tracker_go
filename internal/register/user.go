package register

import (
	"go.mongodb.org/mongo-driver/mongo"

	user_controller "github.com/SamPariat/expenses-tracker/internal/controllers/user"
	"github.com/SamPariat/expenses-tracker/internal/core/services"
	"github.com/SamPariat/expenses-tracker/internal/repositories"
)

func UserRegister(mongoClient *mongo.Client) user_controller.UserController {
	userRepository := repositories.NewUserRepository(mongoClient)
	userService := services.NewUserService(userRepository)
	userController := user_controller.NewUserController(userService)
	return userController
}
