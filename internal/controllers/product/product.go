package product_controller

import (
	"github.com/gin-gonic/gin"

	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type ProductControllerImpl struct {
	productService ports.ProductService
}

func NewProductController(productService ports.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		productService: productService,
	}
}

func (productController *ProductControllerImpl) GetProducts(ctx *gin.Context) {
}

func (productController *ProductControllerImpl) GetProduct(ctx *gin.Context) {
}

func (productController *ProductControllerImpl) CreateProduct(ctx *gin.Context) {
}

func (productController *ProductControllerImpl) UpdateProduct(ctx *gin.Context) {
}

func (productController *ProductControllerImpl) DeleteProduct(ctx *gin.Context) {
}
