package user_controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SamPariat/expenses-tracker/internal/core/ports"
	"github.com/SamPariat/expenses-tracker/utils"
)

type UserControllerImpl struct {
	userService ports.UserService
}

func NewUserController(userService ports.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (userController *UserControllerImpl) GetUsers(ctx *gin.Context) {
	req := new(GetAllUsersRequest)

	if err := req.Validate(ctx); err != nil {
		utils.ReturnBadRequest(ctx, err, "Failed validation for fetching all the users")
		return
	}

	users, err := userController.userService.GetUsers(req.Page, req.Limit)
	if err != nil {
		utils.ReturnInternalServerError(ctx, err, "Something went wrong while fetching all the users")
		return
	}

	utils.ReturnOk(ctx, users, "Succesfully fetched all users")
}

func (userController *UserControllerImpl) GetUser(ctx *gin.Context) {
	req := new(GetUserRequest)

	if err := req.Validate(ctx); err != nil {
		utils.ReturnBadRequest(ctx, err, "Failed validation for fetching the user")
		return
	}

	userId, _ := primitive.ObjectIDFromHex(req.Id)

	user, err := userController.userService.GetUser(userId)
	if err != nil {
		utils.ReturnInternalServerError(ctx, err, "Something went wrong while fetching the user")
		return
	}

	utils.ReturnOk(ctx, user, "Succesfully fetched the user")
}

func (userController *UserControllerImpl) CreateUser(ctx *gin.Context) {
	req := new(CreateUserRequest)

	if err := req.Validate(ctx); err != nil {
		utils.ReturnBadRequest(ctx, err, "Failed validation for creating the user")
		return
	}

	user, err := userController.userService.CreateUser(req.ConvertToUser())
	if err != nil {
		utils.ReturnInternalServerError(ctx, err, "Something went wrong while creating the user")
		return
	}

	utils.ReturnOk(ctx, user, "Succesfully created the user")
}

func (userController *UserControllerImpl) UpdateUser(ctx *gin.Context) {
}

func (userController *UserControllerImpl) DeleteUser(ctx *gin.Context) {
}
