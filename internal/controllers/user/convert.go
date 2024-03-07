package user_controller

import "github.com/SamPariat/expenses-tracker/internal/core/domain"

func (createUserRequest *CreateUserRequest) ConvertToUser() *domain.User {
	return &domain.User{
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
		Income:   createUserRequest.Income,
		ImageUrl: createUserRequest.ImageUrl,
	}
}

func (updateUserRequest *UpdateUserRequest) ConvertToUser() *domain.User {
	return &domain.User{
		Name:     updateUserRequest.Name,
		Email:    updateUserRequest.Email,
		Password: updateUserRequest.Password,
		Income:   updateUserRequest.Income,
		ImageUrl: updateUserRequest.ImageUrl,
	}
}
