package shared

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDb() error {

	// err = godotenv.Load()

	// if err != nil {
	// 	panic(fmt.Sprintf("erro ao carregar variavel de ambiente: %s", err))
	// }

	// dbSslMode := os.Getenv("DB_SSL_MODE")

	var err error

	once.Do(func() {

		if errLoad := godotenv.Load(); errLoad != nil {
			err = fmt.Errorf("erro ao carregar variavel de ambiente: %w", errLoad)
			return
		}
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbPort := os.Getenv("DB_PORT")

		url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, os.Getenv("DB_SSL_MODE"))

		db, err = sql.Open("postgres", url)
		if err != nil {
			return
		}

		err = db.Ping()
		if err != nil {
			return
		}

	})
	// retornando nil, casos nao der erro na conexao
	return err
}

func GetDB() (*sql.DB, error) {
	if db == nil {

		return nil, fmt.Errorf("banco nao inicializado")
	}
	return db, nil
}
