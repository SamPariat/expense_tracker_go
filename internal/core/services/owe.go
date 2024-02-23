package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type OweServiceImpl struct {
	oweRepository ports.OweRepository
}

func NewOweService(OweRepository ports.OweRepository) ports.OweService {
	return &OweServiceImpl{
		oweRepository: OweRepository,
	}
}

func (OweService *OweServiceImpl) GetOwes(page int64, limit int64) ([]domain.Owe, error) {
	return OweService.oweRepository.GetOwes(page, limit)
}

func (OweService *OweServiceImpl) GetOwe(id primitive.ObjectID) (domain.Owe, error) {
	return OweService.oweRepository.GetOwe(id)
}

func (OweService *OweServiceImpl) CreateOwe(owe *domain.Owe) (domain.Owe, error) {
	return OweService.oweRepository.CreateOwe(owe)
}

func (OweService *OweServiceImpl) UpdateOwe(id primitive.ObjectID, owe *domain.Owe) (int64, error) {
	return OweService.oweRepository.UpdateOwe(id, owe)
}

func (OweService *OweServiceImpl) DeleteOwe(id primitive.ObjectID) (int64, error) {
	return OweService.oweRepository.DeleteOwe(id)
}
