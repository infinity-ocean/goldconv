package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/infinity-ocean/goldconv/internal/model"
)

func (r *repo) SelectBalance(id int) (model.Balance, error) {
	conn, err := r.pool.Acquire(context.Background()) // 1
	if err != nil {
		fmt.Println("Pool didn't return connection")
	}
	defer conn.Release()
	selectBalance := sq.Select("balance").From("coins").Where(sq.Eq{"userFK": id})
	sql, _, err := selectBalance.ToSql()
	if err != nil {
		fmt.Println("Squirell failed to make a string for select_balance:", err)
	}
	row := conn.QueryRow(context.Background(), sql)
	if err != nil {
		fmt.Println("pgxpool failed to execute an expression:", err)
	}
	balanceScan := coins{} 
	err = row.Scan(&balanceScan) // 2
	if err != nil {
		fmt.Println("Failed to scan pgx result from select:", err)
	}
	balanceInt := balanceScan.balance
	var balance model.Balance // 3
	balance.Gold = balanceInt / 100
	balance.Silver = uint8((balanceInt % 100) / 10)
	balance.Bronze = uint8(balanceInt % 100)
	return balance, nil
}