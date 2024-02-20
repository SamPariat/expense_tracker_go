package register

import (
	"log"

	product_controller "github.com/SamPariat/expenses-tracker/internal/controllers/product"
	user_controller "github.com/SamPariat/expenses-tracker/internal/controllers/user"
	"github.com/SamPariat/expenses-tracker/internal/pkg/mongo"
)

type AllControllers struct {
	UserController    user_controller.UserController
	ProductController product_controller.ProductController
}

func InitApp() *AllControllers {
	mongoClient, err := mongo.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	userController := UserRegister(mongoClient)
	productController := ProductRegister(mongoClient)

	return &AllControllers{
		UserController:    userController,
		ProductController: productController,
	}
}
