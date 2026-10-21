package database

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		return nil, fmt.Errorf("Database url is empty !!")
	}

	conn, err := sqlx.Connect("pgx", database_url)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect database : %w", err)
	}

	return conn, nil
}
