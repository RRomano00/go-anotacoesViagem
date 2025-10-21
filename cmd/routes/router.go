package routes

import (
	"time"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/handlers"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/repositories"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
		travel.POST("", travelHandler.Create())
		travel.GET("", travelHandler.GetAll())
		travel.GET("/:id", travelHandler.GetTravelByID())
		travel.DELETE("/:id", travelHandler.Delete())
		travel.PATCH("/:id", travelHandler.Update())
		travel.GET("/:id/notes", noteHandler.GetNoteByTravelId())

		travel.POST("/notes", noteHandler.Create())
		travel.DELETE("/notes/:id", noteHandler.DeleteNoteById())
	}

}
