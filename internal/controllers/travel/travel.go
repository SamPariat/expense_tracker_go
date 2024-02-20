package travel_controller

import (
	"github.com/gin-gonic/gin"

	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type TravelControllerImpl struct {
	travelService ports.TravelService
}

func NewTravelController(travelService ports.TravelService) TravelController {
	return &TravelControllerImpl{
		travelService: travelService,
	}
}

func (travelController *TravelControllerImpl) GetTravels(ctx *gin.Context) {

}

func (travelController *TravelControllerImpl) GetTravel(ctx *gin.Context) {

}

func (travelController *TravelControllerImpl) CreateTravel(ctx *gin.Context) {

}

func (travelController *TravelControllerImpl) UpdateTravel(ctx *gin.Context) {

}

func (travelController *TravelControllerImpl) DeleteTravel(ctx *gin.Context) {

}
