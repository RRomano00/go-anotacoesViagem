package repositories

import (
	"database/sql"
	"fmt"

	"github.com/RRomano00/anotacoes_viagem/cmd/internal/models"
	"github.com/RRomano00/anotacoes_viagem/cmd/internal/shared"
)

type TravelRepository struct {
	db *sql.DB
}

// Instanciando repository com suas dependencias (db)
func NewTravelRepository() *TravelRepository {
	db, _ := shared.GetDB()
	return &TravelRepository{db: db} // Passando db para "dentro" do repository
}

func (repo *TravelRepository) Create(travel models.Travel) (models.Travel, error) {
	sql := `INSERT INTO travel (title, start_date, end_date) 
		VALUES ($1,$2,$3) returning id`

	// Tentando inserir valores nos parametros da querry ($), retorna row com id criado
	row := repo.db.QueryRow(sql, travel.Title, travel.StartDate, travel.EndDate)
	if err := row.Err(); err != nil {
		// Retorna objetvo Travel vazio
		return models.Travel{}, err
	}

	// Passa o valor da coluna que foi criada em row para o endereço de memória do id
	var id int
	err := row.Scan(&id)
	if err != nil {
		return models.Travel{}, err
	}

	// Atribui o ID retornado ao objeto original
	travel.Id = id

	return travel, nil

}

func (repo *TravelRepository) GetAll() ([]models.Travel, error) {
	sql := `SELECT id, title, start_date, end_date FROM travel ORDER BY id`

	rows, err := repo.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var travelList []models.Travel
	for rows.Next() {
		var travel models.Travel
		err := rows.Scan(&travel.Id, &travel.Title, &travel.StartDate, &travel.EndDate)
		if err != nil {
			return nil, err
		}

		travelList = append(travelList, travel)
	}

	return travelList, nil
}

func (tr *TravelRepository) GetTravelByID(id int) (models.Travel, error) {
	sql := "SELECT * FROM travel WHERE id = $1"
	resultRow := tr.db.QueryRow(sql, id)

	var travel models.Travel
	err := resultRow.Scan(&travel.Id, &travel.Title, &travel.StartDate, &travel.EndDate)
	if err != nil {
		return models.Travel{}, err
	}

	return travel, nil
}

func (tr *TravelRepository) DeleteTravelById(travelID int) (int64, error) {
	sql := `DELETE FROM travel WHERE id = $1`

	result, err := tr.db.Exec(sql, travelID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (repo *TravelRepository) Update(travel models.UpdateTravelRequest, id int) (int64, error) {

	sql := "UPDATE travel SET "

	var params []any //array que pode receber qualquer tipo de dado
	var idx int = 0

	if travel.Title != "" {
		idx++
		sql += fmt.Sprintf("title = $%d", idx)
		params = append(params, travel.Title)
	}

	if travel.EndDate != nil {
		idx++
		// Só adiciona a virgula se já tiver adicionado o Title
		if idx > 0 {
			sql += ", "
		}
		sql += fmt.Sprintf("end_date = $%d", idx)
		params = append(params, travel.EndDate)

	}

	// Se nenhum campo foi atualizado, nao executa a query
	if idx == 0 {
		return 0, nil // 0 linhas afetadas, sem erro
	}

	idx++
	sql += fmt.Sprintf(" WHERE id = $%d", idx)
	params = append(params, id)

	// result funciona como um resumo da operação que acabou de acontecer
	// err informa se houve erro grave na comunicação ou sintaxe do SQL
	result, err := repo.db.Exec(sql, params...)
	if err != nil {
		return 0, err
	}

	// retornando quantas linhas no bd foram modificados pela query
	return result.RowsAffected()
}
