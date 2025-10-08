package services

import (
	"errors"
	"log/slog"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/models"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/repositories"
)

type TravelService struct {
	travelRepo *repositories.TravelRepository
	noteRepo   *repositories.NoteRepository
}

var ErrTravelNotFound = errors.New("viagem não encontrada")

func NewTravelService(travelRepo *repositories.TravelRepository, noteRepo *repositories.NoteRepository) *TravelService {
	return &TravelService{
		travelRepo: travelRepo,
		noteRepo:   noteRepo,
	}
}

func (ts *TravelService) Create(travel models.Travel) (models.Travel, error) {
	return ts.travelRepo.Create(travel)
}

func (ts *TravelService) GetAll() ([]models.Travel, error) {
	return ts.travelRepo.GetAll()
}

func (ts *TravelService) GetTravelByID(id int) (models.Travel, error) {
	return ts.travelRepo.GetTravelByID(id)
}

func (ts *TravelService) DeleteTravelAndNotes(travelID int) error {

	// Deleta as anotações associadas com a viagem
	err := ts.noteRepo.DeleteByTravelID(travelID)
	if err != nil {
		slog.Error("erro ao deletar anotações associadas", "travelId", travelID, "error", err)
		return err
	}

	// Delete a viagem e captura o numero de linhas afetadas
	rowsAffected, err := ts.travelRepo.DeleteTravelById(travelID)
	if err != nil {
		slog.Error("erro ao deletar a viagem", "travelId", travelID, "error", err)
		return err
	}

	// se 0 linhas afetadas, significa que o ID não foi encontrado
	if rowsAffected == 0 {
		return ErrTravelNotFound
	}

	return nil
}

func (ts *TravelService) Update(travel models.UpdateTravelRequest, travelId int) error {
	rowsAffected, err := ts.travelRepo.Update(travel, travelId)
	if err != nil {
		return err
	}

	// se nenhuma linha foi afetada, o ID não existe
	if rowsAffected == 0 {
		return ErrTravelNotFound
	}

	return nil
}
