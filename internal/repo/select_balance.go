package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/infinity-ocean/goldconv/internal/model"
)

func (r *repo) SelectBalance(id int) (model.Balance, error) {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("Pool didn't return connection")
	}
	defer conn.Conn().Close(context.Background())
	selectBalance := sq.Select("balance").From("coins").Where(sq.Eq{"userFK": id})
	sql, _, err := selectBalance.ToSql()
	if err != nil {
		fmt.Println("Squirell failed to make a string for select_balance:", err)
	}
	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println("pgxpool failed to execute an expression:", err)
	}
	var balance model.Balance
	err = rows.Scan(&balance)
	if err != nil {
		fmt.Println("Failed to scan pgx result from select:", err)
	}
	return balance, nil
}

