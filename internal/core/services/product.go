package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type ProductServiceImpl struct {
	productRepository ports.ProductRepository
}

func NewProductService(productRepository ports.ProductRepository) ports.ProductService {
	return &ProductServiceImpl{
		productRepository: productRepository,
	}
}

func (productService *ProductServiceImpl) GetProducts(page int64, limit int64) ([]domain.Product, error) {
	return productService.productRepository.GetProducts(page, limit)
}

func (productService *ProductServiceImpl) GetProduct(id primitive.ObjectID) (domain.Product, error) {
	return productService.productRepository.GetProduct(id)
}

func (productService *ProductServiceImpl) CreateProduct(user *domain.Product) (domain.Product, error) {
	return productService.productRepository.CreateProduct(user)
}

func (productService *ProductServiceImpl) UpdateProduct(id primitive.ObjectID) (domain.Product, error) {
	return productService.productRepository.UpdateProduct(id)
}

func (productService *ProductServiceImpl) DeleteProduct(id primitive.ObjectID) (domain.Product, error) {
	return productService.productRepository.DeleteProduct(id)
}
