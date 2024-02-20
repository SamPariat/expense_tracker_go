package register

import (
	"go.mongodb.org/mongo-driver/mongo"

	product_controller "github.com/SamPariat/expenses-tracker/internal/controllers/product"
	"github.com/SamPariat/expenses-tracker/internal/core/services"
	"github.com/SamPariat/expenses-tracker/internal/repositories"
)

func ProductRegister(mongoClient *mongo.Client) product_controller.ProductController {
	productRepository := repositories.NewProductRepository(mongoClient)
	productService := services.NewProductService(productRepository)
	productController := product_controller.NewProductController(productService)
	return productController
}
