package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type UserServiceImpl struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (userService *UserServiceImpl) GetUsers(page int64, limit int64) ([]domain.User, error) {
	return userService.userRepository.GetUsers(page, limit)
}

func (userService *UserServiceImpl) GetUser(id primitive.ObjectID) (domain.User, error) {
	return userService.userRepository.GetUser(id)
}

func (userService *UserServiceImpl) CreateUser(user *domain.User) (domain.User, error) {
	return userService.userRepository.CreateUser(user)
}

func (userService *UserServiceImpl) UpdateUser(id primitive.ObjectID) (domain.User, error) {
	return userService.userRepository.UpdateUser(id)
}

func (userService *UserServiceImpl) DeleteUser(id primitive.ObjectID) (domain.User, error) {
	return userService.userRepository.DeleteUser(id)
}
