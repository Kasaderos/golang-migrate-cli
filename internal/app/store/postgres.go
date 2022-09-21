package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"git.dar.kz/forte-market/migrations/internal/app/config"
)

func InitPostgresConnection(cfg config.PostgresConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
