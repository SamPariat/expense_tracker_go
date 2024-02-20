package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type TravelServiceImpl struct {
	travelRepository ports.TravelRepository
}

func NewTravelService(travelRepository ports.TravelRepository) ports.TravelService {
	return &TravelServiceImpl{
		travelRepository: travelRepository,
	}
}

func (travelService *TravelServiceImpl) GetTravels(page int64, limit int64) ([]domain.Travel, error) {
	return travelService.travelRepository.GetTravels(page, limit)
}

func (travelService *TravelServiceImpl) GetTravel(id primitive.ObjectID) (domain.Travel, error) {
	return travelService.travelRepository.GetTravel(id)
}

func (travelService *TravelServiceImpl) CreateTravel(travel *domain.Travel) (domain.Travel, error) {
	return travelService.travelRepository.CreateTravel(travel)
}

func (travelService *TravelServiceImpl) UpdateTravel(id primitive.ObjectID) (domain.Travel, error) {
	return travelService.travelRepository.UpdateTravel(id)
}

func (travelService *TravelServiceImpl) DeleteTravel(id primitive.ObjectID) (domain.Travel, error) {
	return travelService.travelRepository.DeleteTravel(id)
}
