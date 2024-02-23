package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type TravelRepositoryImpl struct {
	mongoClient *mongo.Client
}

func NewTravelRepository(mongoClient *mongo.Client) ports.TravelRepository {
	return &TravelRepositoryImpl{
		mongoClient: mongoClient,
	}
}

func (travelRepository *TravelRepositoryImpl) GetTravels(page int64, limit int64) ([]domain.Travel, error) {
	return []domain.Travel{}, nil
}

func (travelRepository *TravelRepositoryImpl) GetTravel(id primitive.ObjectID) (domain.Travel, error) {
	return domain.Travel{}, nil
}

func (travelRepository *TravelRepositoryImpl) CreateTravel(travel *domain.Travel) (domain.Travel, error) {
	return domain.Travel{}, nil
}

func (travelRepository *TravelRepositoryImpl) UpdateTravel(id primitive.ObjectID, travel *domain.Travel) (int64, error) {
	return 0, nil
}

func (travelRepository *TravelRepositoryImpl) DeleteTravel(id primitive.ObjectID) (int64, error) {
	return 0, nil
}
