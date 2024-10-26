package postgresRepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

var DB *sql.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// InitDB inicializa a conexão com o banco de dados
func InitDB() (*sql.DB, error) {
	config := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Verifica a conexão
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}

func InitContextAndDB() (context.Context, context.CancelFunc, *sql.DB, error) {
	// TODO:
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)

	dbConnection, err := InitDB()
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}

	return ctx, cancel, dbConnection, nil
}

// CloseDB fecha a conexão com o banco de dados
func CloseDB() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}
}
