package repositories

import (
	"database/sql"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/models"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/shared"
)

type NoteRepository struct {
	db *sql.DB
}

func NewNoteRepository() *NoteRepository {
	db, _ := shared.GetDB()
	return &NoteRepository{db: db}
}

func (repo *NoteRepository) Create(note models.Note) error {
	sql := `INSERT INTO note (content, travel_id, created_at) VALUES ($1,$2, $3)`

	row := repo.db.QueryRow(sql, note.Content, note.Travel_Id, note.Created_at)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}

func (nr *NoteRepository) GetNoteByTravelID(travelID int) ([]models.NoteTravel, error) {

	sql := "SELECT travel.title, note.* FROM note INNER JOIN travel ON travel.id = note.travel_id WHERE travel.id = $1"

	// passando parametro travelID para a querry
	rows, err := nr.db.Query(sql, travelID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// slice para armazenar os resultas
	var notes []models.NoteTravel

	// Itera sobre os resultados encontrados pelo banco e preenche struct
	for rows.Next() {
		var note models.NoteTravel

		// Escaneia a linha atual para dentro da struct
		err := rows.Scan(
			&note.TravelName,
			&note.Id,
			&note.Travel_Id,
			&note.Content,
			&note.Created_at,
		)
		if err != nil {
			return nil, err // erro se a leitura falhar
		}
		// Adiciona ao alice as structs preenchidas
		notes = append(notes, note)
	}

	return notes, nil
}

func (nr *NoteRepository) DeleteByTravelID(travelID int) error {
	sql := `DELETE FROM note WHERE travel_id = $1`

	// Exec nao retorna nenhuma linha (INSERT, UPDATE e DELETE)
	_, err := nr.db.Exec(sql, travelID)

	if err != nil {
		return err
	}

	return nil
}

func (nr *NoteRepository) DeleteByID(id int) error {
	sql := `DELETE FROM note WHERE id = $1`

	// Exec nao retorna nenhuma linha (INSERT, UPDATE e DELETE)
	_, err := nr.db.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}
