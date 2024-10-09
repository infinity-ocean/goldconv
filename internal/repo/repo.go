package repo

import (
	"context"
	"fmt"
	"github.com/infinity-ocean/goldconv/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct{
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) *repo {
	return &repo{pool: pool}
}

func (r *repo) InsertAccount(small model.AccountSmall) error {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()
    sql := fmt.Sprintf(
        `INSERT INTO accounts (username, email, password, balance) 
		VALUES ('%s', '%s', '%s', '%s');`,
		small.Username,
		small.Email, 
		small.Password, 
		small.Balance,
	)
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("can't insert account: %w", err)
   	}
	return nil
}