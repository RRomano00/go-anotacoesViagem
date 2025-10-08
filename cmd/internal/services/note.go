package services

import (
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/models"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/repositories"
)

type NoteService struct {
	repo *repositories.NoteRepository
}

func NewNoteService(repo *repositories.NoteRepository) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (ns *NoteService) Create(note models.Note) error {
	return ns.repo.Create(note)
}

func (ns *NoteService) GetNoteByTravelID(travelID int) ([]models.NoteTravel, error) {
	return ns.repo.GetNoteByTravelID(travelID)
}
