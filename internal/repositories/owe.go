package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type OweRepositoryImpl struct {
	mongoClient *mongo.Client
}

func NewOweRepository(mongoClient *mongo.Client) ports.OweRepository {
	return &OweRepositoryImpl{
		mongoClient: mongoClient,
	}
}

func (oweRepositoryImpl *OweRepositoryImpl) GetOwes(page int64, limit int64) ([]domain.Owe, error) {
	return []domain.Owe{}, nil
}

func (oweRepositoryImpl *OweRepositoryImpl) GetOwe(id primitive.ObjectID) (domain.Owe, error) {
	return domain.Owe{}, nil
}

func (oweRepositoryImpl *OweRepositoryImpl) CreateOwe(owe *domain.Owe) (domain.Owe, error) {
	return domain.Owe{}, nil
}

func (oweRepositoryImpl *OweRepositoryImpl) UpdateOwe(id primitive.ObjectID, owe *domain.Owe) (int64, error) {
	return 0, nil
}

func (oweRepositoryImpl *OweRepositoryImpl) DeleteOwe(id primitive.ObjectID) (int64, error) {
	return 0, nil
}
