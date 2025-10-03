package services

import (
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/models"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/repositories"
)

type TravelService struct {
	repo *repositories.TravelRepository
}

func NewTravelService(repo *repositories.TravelRepository) *TravelService {
	return &TravelService{
		repo: repo,
	}
}

func (service *TravelService) Create(travel models.Travel) (models.Travel, error) {
	return service.repo.Create(travel)
}
