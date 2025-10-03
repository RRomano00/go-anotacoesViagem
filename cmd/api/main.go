package main

import (
	"fmt"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/shared"
	"github.com/RRomano00/anotacoes_viagem/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	err := shared.InitDb()
	if err != nil {
		panic(fmt.Sprintf("erro ao inicializar banco:%s", err))
	}

	// Cria roteador do Gin
	router := gin.Default()

	routes.RegisterRoutes(router)

	fmt.Println("Listening... (8080)")
	router.Run(":8080")
}
