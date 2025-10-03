package routes

import (
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/handlers"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/repositories"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// ---- Injeção de Dependências ----
	// 1. Criando repository
	travelRepository := repositories.NewTravelRepository()

	// 2. Injetando o Repo no Service
	travelService := services.NewTravelService(travelRepository)

	// 3. Injetando o Service no Handler
	travelHandler := handlers.NewTravelHandler(travelService)
	// --------------------

	// ---- Definição das rotas ----
	// Agrupa rotas relacionadas a 'travel'
	api := router.Group("/travels")
	{
		api.POST("/", travelHandler.Create())
	}

}
