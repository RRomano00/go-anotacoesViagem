package main

import (
	"fmt"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/shared"
)

func main() {

	err := shared.InitDb()
	if err != nil {
		panic(fmt.Sprintf("erro ao inicializar banco:%s", err))
	}


	
}
