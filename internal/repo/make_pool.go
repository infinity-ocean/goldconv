package repo

import (
	"context"
	"fmt"
	"github.com/infinity-ocean/goldconv/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func MakePool(config config.Config) (*pgxpool.Pool, error) {
	err := godotenv.Load("infra.env")
	if err != nil {
		return &pgxpool.Pool{}, fmt.Errorf("failed to load env: %w", err)
	}
	// const DB = "postgres://postgres:12345@localhost:5432/postgres"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.PGUSER,
		config.PGPASSWORD,
		config.PGHost,
		config.PGPORT,
		config.PGDB,
		config.PGSSL,
	)

	// Create the connection pool
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}
	return dbpool, nil
}
