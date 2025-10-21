package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/models"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/services"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	service *services.NoteService
}

func NewNoteHandler(noteService *services.NoteService) *NoteHandler {
	return &NoteHandler{
		service: noteService,
	}
}

func (nh *NoteHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pega o conteudo da nota a partir do body do JSON
		var req models.CreateNoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		note := models.Note{
			Content:    req.Content,
			Travel_Id:  req.TravelID,
			Created_at: time.Now(),
		}

		// Passa o note preenchido para a camada Service
		err := nh.service.Create(note) // Service processa a lógica de negóicos (salva no bd, etc)
		if err != nil {
			slog.Error("Erro ao criar anotação", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Status 201 (Note criado!). Passa o note criado de volta no corpo da resposta
		c.JSON(http.StatusCreated, "note criado com sucesso!")
	}
}

func (nh *NoteHandler) GetNoteByTravelId() gin.HandlerFunc {
	return func(c *gin.Context) {
		travelId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		note, err := nh.service.GetNoteByTravelID(travelId)

		if err != nil {
			slog.Error("Erro ao buscar anotações", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// converte o slice que retornou (handler <- service <- repo), em uma string no formato JSON
		// envia esse JSON de volta no corpo da response HTTP
		c.IndentedJSON(http.StatusOK, note)
	}
}

func (nh *NoteHandler) DeleteNoteById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		err = nh.service.DeleteTravelById(id)

		if err != nil {
			slog.Error("Erro ao deletar anotaçõe", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// converte o slice que retornou (handler <- service <- repo), em uma string no formato JSON
		// envia esse JSON de volta no corpo da response HTTP
		c.IndentedJSON(http.StatusOK, "Deletado com sucesso!")
	}
}
