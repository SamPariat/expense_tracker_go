package product_controller

import (
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetProducts(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}
