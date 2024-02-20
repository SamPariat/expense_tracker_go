package travel_controller

import "github.com/gin-gonic/gin"

type TravelController interface {
	GetTravels(ctx *gin.Context)
	GetTravel(ctx *gin.Context)
	CreateTravel(ctx *gin.Context)
	UpdateTravel(ctx *gin.Context)
	DeleteTravel(ctx *gin.Context)
}
