package database

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	mu   sync.Mutex
	pool *sqlx.DB
)

func ConnectDB() (*sqlx.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	if pool != nil {
		return pool, nil
	}
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, errors.New("DATABASE_URL is empty")
	}
	conn, err := sqlx.Connect("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(30 * time.Minute)

	pool = conn
	return pool, nil
}
