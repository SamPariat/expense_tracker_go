package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type UserRepositoryImpl struct {
	mongoClient *mongo.Client
}

func NewUserRepository(mongoClient *mongo.Client) ports.UserRepository {
	return &UserRepositoryImpl{
		mongoClient: mongoClient,
	}
}

func (userRepository *UserRepositoryImpl) GetUsers(page int64, limit int64) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (userRepository *UserRepositoryImpl) GetUser(id primitive.ObjectID) (domain.User, error) {
	return domain.User{}, nil
}

func (userRepository *UserRepositoryImpl) CreateUser(user *domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (userRepository *UserRepositoryImpl) UpdateUser(id primitive.ObjectID) (domain.User, error) {
	return domain.User{}, nil
}

func (userRepository *UserRepositoryImpl) DeleteUser(id primitive.ObjectID) (domain.User, error) {
	return domain.User{}, nil
}
