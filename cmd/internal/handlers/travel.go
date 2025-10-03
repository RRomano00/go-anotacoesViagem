package handlers

import (
	"net/http"

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

func (h *TravelHandler) Create() gin.HandlerFunc {

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
		createdTravel, err := h.service.Create(travel) // Service processa a lógica de negóicos (salva no bd, etc)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Status 201 (Travel criado!). Passa a viagem criada de volta no corpo da resposta
		c.JSON(http.StatusCreated, createdTravel)
	}

}
