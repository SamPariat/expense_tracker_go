package user_controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}
