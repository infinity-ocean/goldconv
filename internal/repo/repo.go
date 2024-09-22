package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct{
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) *repo {
	return &repo{pool: pool}
}