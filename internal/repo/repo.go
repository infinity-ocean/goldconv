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

func (r *repo) InsertAccount(id int, small model.AccountSmall) error {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()
	var exists bool
	err = conn.QueryRow(context.Background(),"SELECT EXISTS(SELECT 1 FROM accounts WHERE id = $1)", id).Scan(&exists)
	if err != nil {
	    return fmt.Errorf("failed select account: %w", err)
	}
	if exists {
		return fmt.Errorf("account already exists: %w", err)
	}
    sql := fmt.Sprintf(
        `INSERT INTO accounts (id, username, email, password, balance) 
		VALUES (%d, %s, %s, %s, %s) 
		ON CONFLICT (userFK) 
		DO UPDATE SET balance = EXCLUDED.balance;`,
        id,
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