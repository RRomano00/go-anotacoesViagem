package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/models"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/services"
	"github.com/gin-gonic/gin"
)

type TravelHandler struct {
	service *services.TravelService
}

// Construtor do Handler que recebe um travelService como parametro
func NewTravelHandler(travelService *services.TravelService) *TravelHandler {
	// Pega o parametro (travelService) e atribui ao handler (propriedade service)
	return &TravelHandler{
		service: travelService,
	}
}

func (th *TravelHandler) Create() gin.HandlerFunc {

	// gin.Context carrega todas as info da requisição e da resposta
	return func(c *gin.Context) {
		// Variavel que recebe os dados JSON
		travel := models.Travel{}

		//  Tenta ligar (bind) o JSON do corpo da req à stuct
		err := c.ShouldBindJSON(&travel)
		if err != nil { // Caso JSON inválido ou os tipos nao baterem
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Passa o travel preenchido para a camada Service
		createdTravel, err := th.service.Create(travel) // Service processa a lógica de negóicos (salva no bd, etc)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Status 201 (Travel criado!). Passa a viagem criada de volta no corpo da resposta
		c.JSON(http.StatusCreated, createdTravel)
	}
}

func (th *TravelHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		travelList, err := th.service.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, travelList)
	}
}

func (th *TravelHandler) GetTravelByID() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id")) // Converte str pra int

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		travel, err := th.service.GetTravelByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, travel)
	}
}

func (th *TravelHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		travelId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		err = th.service.DeleteTravelAndNotes(travelId)
		// Verificando se o erro é do tipo ErrTravelNotFound
		if errors.Is(err, services.ErrTravelNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Se for qualquer outro tipo de erro
		if err != nil {
			slog.Error("Erro ao deletar viagem", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
			return
		}

		// Nao
		c.Status(http.StatusNoContent)
	}
}

func (th *TravelHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		travelId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		// req vai receber dados do JSON
		var req models.UpdateTravelRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Corpo da requisição inválido"})
			return
		}

		// manda para o service os dados recebidos
		err = th.service.Update(req, travelId)

		// Verificando se o erro é do tipo ErrTravelNotFound
		if errors.Is(err, services.ErrTravelNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Se for qualquer outro tipo de erro
		if err != nil {
			slog.Error("Erro ao atualizar viagem", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
			return
		}

		// tudo deu certo, 200 OK
		c.JSON(http.StatusOK, gin.H{"message": "Viagem atualizada com sucesso!"})

	}
}
