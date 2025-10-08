package routes

import (
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/handlers"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/repositories"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// ---- Injeção de Dependências ----
	travelRepo := repositories.NewTravelRepository()
	noteRepo := repositories.NewNoteRepository()

	noteService := services.NewNoteService(noteRepo)
	noteHandler := handlers.NewNoteHandler(noteService)

	travelService := services.NewTravelService(travelRepo, noteRepo)
	travelHandler := handlers.NewTravelHandler(travelService)

	// ----------------------------------------

	// ---- Definição das rotas ----
	// Agrupa rotas relacionadas a 'travel'
	travel := router.Group("/travels")
	{
		travel.POST("/", travelHandler.Create())
		travel.GET("/", travelHandler.GetAll())
		travel.GET("/:id", travelHandler.GetTravelByID())
		travel.DELETE("/:id", travelHandler.Delete())

		travel.PATCH("/:id", travelHandler.Update())

		travel.POST("/notes", noteHandler.Create())
		travel.GET("/:id/notes", noteHandler.GetNoteByTravelId())
	}

}
