package ports

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
)

type UserRepository interface {
	GetUsers(page int64, limit int64) ([]domain.User, error)
	GetUser(id primitive.ObjectID) (domain.User, error)
	CreateUser(user *domain.User) (domain.User, error)
	UpdateUser(id primitive.ObjectID) (domain.User, error)
	DeleteUser(id primitive.ObjectID) (domain.User, error)
}

type ProductRepository interface {
	GetProducts(page int64, limit int64) ([]domain.Product, error)
	GetProduct(id primitive.ObjectID) (domain.Product, error)
	CreateProduct(product *domain.Product) (domain.Product, error)
	UpdateProduct(id primitive.ObjectID) (domain.Product, error)
	DeleteProduct(id primitive.ObjectID) (domain.Product, error)
}

type TravelRepository interface {
	GetTravels(page int64, limit int64) ([]domain.Travel, error)
	GetTravel(id primitive.ObjectID) (domain.Travel, error)
	CreateTravel(travel *domain.Travel) (domain.Travel, error)
	UpdateTravel(id primitive.ObjectID) (domain.Travel, error)
	DeleteTravel(id primitive.ObjectID) (domain.Travel, error)
}
