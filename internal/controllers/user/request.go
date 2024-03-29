package user_controller

import (
	"github.com/gin-gonic/gin"
)

type GetAllUsersRequest struct {
	Page  int64 `form:"page" validate:"required,gt=0"`
	Limit int64 `form:"limit" validate:"required,gt=0"`
}

type GetUserRequest struct {
	Id string `uri:"id" validate:"required,mongodb"`
}

type CreateUserRequest struct {
	Name     string  `json:"name" bson:"name" validate:"required,min=3"`
	Email    string  `json:"email" bson:"email" validate:"required,email"`
	Password string  `json:"password" bson:"password" validate:"required,min=7"`
	Income   float64 `json:"income" bson:"income" validate:"required,gt=0"`
	ImageUrl *string `json:"imageUrl" bson:"imageUrl" validate:"omitempty,url"`
}

type UpdateUserRequest struct {
	Name     string  `json:"name" bson:"name" validate:"min=3"`
	Email    string  `json:"email" bson:"email" validate:"email"`
	Password string  `json:"password" bson:"password" validate:"min=7"`
	Income   float64 `json:"income" bson:"income" validate:"gt=0"`
	ImageUrl *string `json:"imageUrl" bson:"imageUrl" validate:"omitempty,url"`
}

func (getAllUsersRequest *GetAllUsersRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindQuery(getAllUsersRequest)
}

func (getUserRequest *GetUserRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindUri(getUserRequest)
}

func (createUserRequest *CreateUserRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(createUserRequest)
}

func (updateUserRequest *UpdateUserRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&updateUserRequest)
}
