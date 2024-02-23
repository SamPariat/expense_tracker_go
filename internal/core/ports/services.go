package ports

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
)

type UserService interface {
	GetUsers(page int64, limit int64) ([]domain.User, error)
	GetUser(id primitive.ObjectID) (domain.User, error)
	CreateUser(user *domain.User) (domain.User, error)
	UpdateUser(id primitive.ObjectID, user *domain.User) (int64, error)
	DeleteUser(id primitive.ObjectID) (int64, error)
}

type ProductService interface {
	GetProducts(page int64, limit int64) ([]domain.Product, error)
	GetProduct(id primitive.ObjectID) (domain.Product, error)
	CreateProduct(product *domain.Product) (domain.Product, error)
	UpdateProduct(id primitive.ObjectID, product *domain.Product) (int64, error)
	DeleteProduct(id primitive.ObjectID) (int64, error)
}

type TravelService interface {
	GetTravels(page int64, limit int64) ([]domain.Travel, error)
	GetTravel(id primitive.ObjectID) (domain.Travel, error)
	CreateTravel(travel *domain.Travel) (domain.Travel, error)
	UpdateTravel(id primitive.ObjectID, travel *domain.Travel) (int64, error)
	DeleteTravel(id primitive.ObjectID) (int64, error)
}

type OweService interface {
	GetOwes(page int64, limit int64) ([]domain.Owe, error)
	GetOwe(id primitive.ObjectID) (domain.Owe, error)
	CreateOwe(owe *domain.Owe) (domain.Owe, error)
	UpdateOwe(id primitive.ObjectID, owe *domain.Owe) (int64, error)
	DeleteOwe(id primitive.ObjectID) (int64, error)
}
