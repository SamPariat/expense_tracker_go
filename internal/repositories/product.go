package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type ProductRepositoryImpl struct {
	mongoClient *mongo.Client
}

func NewProductRepository(mongoClient *mongo.Client) ports.ProductRepository {
	return &ProductRepositoryImpl{
		mongoClient: mongoClient,
	}
}

func (productRepository *ProductRepositoryImpl) GetProducts(page int64, limit int64) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (productRepository *ProductRepositoryImpl) GetProduct(id primitive.ObjectID) (domain.Product, error) {
	return domain.Product{}, nil
}

func (productRepository *ProductRepositoryImpl) CreateProduct(product *domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}

func (productRepository *ProductRepositoryImpl) UpdateProduct(id primitive.ObjectID) (domain.Product, error) {
	return domain.Product{}, nil
}

func (productRepository *ProductRepositoryImpl) DeleteProduct(id primitive.ObjectID) (domain.Product, error) {
	return domain.Product{}, nil
}
