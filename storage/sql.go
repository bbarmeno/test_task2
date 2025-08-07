package storage

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresConfig struct {
	DatabaseDSN string `env:"DATABASE_DSN" envDefault:"root:root@tcp(127.0.0.1:5432)/test"`
}

func NewSql(cfg PostgresConfig) (*sqlx.DB, error) {
	if cfg.DatabaseDSN == "" {
		return nil, fmt.Errorf("config is nil")
	}

	db, err := sqlx.Open("pgx", fmt.Sprintf("%s?sslmode=disable", cfg.DatabaseDSN))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	return db, nil
}
