package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnOk(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    data,
		"error":   nil,
		"message": message,
	})
}

func ReturnCreated(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  201,
		"data":    data,
		"error":   nil,
		"message": message,
	})
}

func ReturnBadRequest(ctx *gin.Context, err error, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  400,
		"data":    nil,
		"error":   err,
		"message": message,
	})
}

func ReturnUnauthorized(ctx *gin.Context, err error, message string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"status":  401,
		"data":    nil,
		"error":   err,
		"message": message,
	})
}

func ReturnNotFound(ctx *gin.Context, err error, message string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"data":    nil,
		"error":   err,
		"message": message,
	})
}

func ReturnInternalServerError(ctx *gin.Context, err error, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"data":    nil,
		"error":   err,
		"message": message,
	})
}
