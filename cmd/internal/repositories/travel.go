package repositories

import (
	"database/sql"

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
	sql := `INSERT INTO TRAVEL (title, start_date, end_date) 
		values ($1,$2,$3) returning id`

	// Tentando inserir valores nos parametros da querry ($), retorna row com id criado
	row := repo.db.QueryRow(sql, travel.Title, travel.Start_date, travel.End_date)
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
