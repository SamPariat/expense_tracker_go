package register

import (
	"log"

	user_controller "github.com/SamPariat/expenses-tracker/internal/controllers/user"
	"github.com/SamPariat/expenses-tracker/internal/pkg/mongo"
)

type AllControllers struct {
	UserController user_controller.UserController
}

func InitApp() *AllControllers {
	mongoClient, err := mongo.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	userController := UserRegister(mongoClient)

	return &AllControllers{
		UserController: userController,
	}
}
