package main

import (
	"fmt"
	"log"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/shared"
	"github.com/RRomano00/anotacoes_viagem/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	err := shared.InitDb()
	if err != nil {
		panic(fmt.Sprintf("erro ao inicializar banco:%s", err))
	}

	router := gin.Default()

	routes.RegisterRoutes(router)

	fmt.Println("Listening... (8080)")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
