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
