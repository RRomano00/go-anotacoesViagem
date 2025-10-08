package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/models"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/shared"
)

const (
	CREATE_TRAVEL = "1"
	LIST_TRAVEL   = "2"
	CREATE_NOTE   = "3"
	NOTE_BYID     = "4"
	DELETE_TRAVEL = "5"
	UPDATE_TRAVEL = "6"
	EXIT          = "0"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	exit := false

	for !exit {

		fmt.Println("------- Diario Digital -------")

		fmt.Println(CREATE_TRAVEL, " - Criar viagem")
		fmt.Println(LIST_TRAVEL, " - Listar viagens")
		fmt.Println(CREATE_NOTE, " - Criar anotação")
		fmt.Println(NOTE_BYID, " - Pesquisar anotações por travelID")
		fmt.Println(DELETE_TRAVEL, " - Deletar uma viagem e sua anotações")
		fmt.Println(UPDATE_TRAVEL, " - Atualizar uma viagem")
		fmt.Println(EXIT, " - Sair...")

		optionInput, _ := reader.ReadString('\n')

		if strings.TrimSpace(optionInput) == EXIT {
			exit = true
			fmt.Println("Saindo...")
			continue
		}

		if strings.TrimSpace(optionInput) == CREATE_TRAVEL {
			CreateTravelTerminal(reader)
			continue
		}

		if strings.TrimSpace(optionInput) == LIST_TRAVEL {
			GetAllTravelTerminal(reader)
			continue
		}

		if strings.TrimSpace(optionInput) == CREATE_NOTE {
			CreateNoteTerminal(reader)
			continue
		}
		if strings.TrimSpace(optionInput) == NOTE_BYID {
			GetNoteByTravelIDTerminal(reader)
			continue
		}
		if strings.TrimSpace(optionInput) == DELETE_TRAVEL {
			DeleteTravelById(reader)
			continue
		}
		if strings.TrimSpace(optionInput) == UPDATE_TRAVEL {
			UpdateTravelTerminal(reader)
			continue
		}

	}

}

func CreateTravelTerminal(reader *bufio.Reader) {

	fmt.Print("Titulo da viagem: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Data de inicio: ")
	startDate, _ := reader.ReadString('\n')

	startDateTime, err := time.Parse(shared.DD_MM_YYYY, strings.TrimSpace(startDate))

	if err != nil {
		slog.Error(err.Error())
		return
	}

	travel := models.Travel{
		Title:     strings.TrimSpace(title),
		StartDate: startDateTime,
	}

	p, err := json.Marshal(travel)
	if err != nil {
		panic(err)
	}

	_, err = http.Post("http://localhost:8080/travels", "Content-Type: application/json", bytes.NewReader(p))

	if err != nil {
		fmt.Println("erro ao criar viagem ", err)
		return
	}

	fmt.Println("Viagem criada com sucesso!")
}

func CreateNoteTerminal(reader *bufio.Reader) {
	fmt.Print("Descrição: ")
	content, _ := reader.ReadString('\n')

	fmt.Print("Id da viagem: ")
	id, _ := reader.ReadString('\n')
	idInt, err := strconv.Atoi(strings.TrimSpace(id))

	if err != nil {
		fmt.Println(http.StatusBadRequest)
		return
	}

	note := models.CreateNoteRequest{
		Content:  content,
		TravelID: idInt,
	}

	p, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/travels/notes", "Content-Type: application/json", bytes.NewReader(p))

	if err != nil {
		fmt.Println("erro ao criar anotação ", err)
		return
	}

	if resp.StatusCode != 201 {
		fmt.Println("erro!!")
		return
	}

	fmt.Println("Anotação criada com sucesso!")
}

func GetAllTravelTerminal(reader *bufio.Reader) {

	resp, err := http.Get("http://localhost:8080/travels")
	if err != nil {
		fmt.Println("🚨 Erro ao tentar conectar com a API. 🚨")
		fmt.Printf("Detalhe do erro: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		var travels []models.Travel

		if err := json.NewDecoder(resp.Body).Decode(&travels); err != nil {
			fmt.Println(err.Error())
		}

		if len(travels) > 0 {
			fmt.Println("\n--- Lista de Viagens ---")
			for _, travel := range travels {
				endDateStr := "Não definida" // caso o usuario nao tenha colocado uma data final

				if travel.EndDate != nil {
					endDateStr = travel.EndDate.Format(shared.DD_MM_YYYY)
				}

				formatedStartDate := travel.StartDate.Format(shared.DD_MM_YYYY)

				fmt.Printf("Viagem ID: %d | Titulo: %s | Data de inicio: %s | Data fim: %s\n", travel.Id, travel.Title, formatedStartDate, endDateStr)
			}
			fmt.Println("---------------------")
		} else {
			fmt.Println("\nNenhuma viagem encontrada!")
		}

	} else {
		fmt.Printf("\nErro da API - Status: %s\n\n", resp.Status)
	}
}

func GetNoteByTravelIDTerminal(reader *bufio.Reader) {

	fmt.Println("Insira o id da viagem: ")

	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	resp, err := http.Get("http://localhost:8080/travels/" + id + "/notes")

	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		var notes []models.NoteTravel

		// Converte o JSON da resposta para structs do Go
		if err := json.NewDecoder(resp.Body).Decode(&notes); err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(notes) > 0 {
			for _, note := range notes {
				formatedDate := note.Created_at.Format(shared.DD_MM_YYYY)
				fmt.Printf("%v - Anotação | %s \n Conteúdo: %s Data de criação: %s \n", note.Id, note.TravelName, note.Content, formatedDate)
				fmt.Println("")
			}
		} else {
			fmt.Println("Nenhuma viagem encontrada!")
		}
	}
}

func DeleteTravelById(reader *bufio.Reader) {
	fmt.Println("Insira o id da viagem a ser removida: ")

	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	url := "http://localhost:8080/travels/" + id

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		fmt.Println("Erro ao criar requisição: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar requisição: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Viagem e anotações excluídas com sucesso!")
	} else {
		fmt.Printf("Erro ao excluir viagem. Status %s\n", resp.Status)
	}
}

func UpdateTravelTerminal(reader *bufio.Reader) {
	fmt.Println("Insira o id da viagem a ser atualizada: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Println("Novo título (em branco para nao alterar): ")
	newTitle, _ := reader.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)

	fmt.Println("Nova data final (DD/MM/AAAA, deixe em branco para não alterar): ")
	newEndDateStr, _ := reader.ReadString('\n')
	newEndDateStr = strings.TrimSpace(newEndDateStr)

	updateReq := models.UpdateTravelRequest{
		Title: newTitle,
	}

	if newEndDateStr != "" {
		parsedDate, err := time.Parse(shared.DD_MM_YYYY, newEndDateStr)
		if err != nil {
			slog.Error("Formato da data inválido. Utilize DD/MM/AAAA.", "error", err)
			return
		}
		updateReq.EndDate = &parsedDate
	}

	jsonData, err := json.Marshal(updateReq)
	if err != nil {
		fmt.Println("Erro ao criar JSON", "error", err)
		return
	}

	url := "http://localhost:8080/travels/" + id
	// NewBuffer le data existente em um slice
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erro ao criar requisição: ", err)
		return
	}

	// adicionando cabeçalho para que Gin entenda que é JSON
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar requisição: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Viagem alterada com sucesso!")
	} else {
		fmt.Printf("Erro ao alterar viagem. Status %s\n", resp.Status)
	}
}
