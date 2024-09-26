package repo

import (
	"context"
	"fmt"

	// sq "github.com/Masterminds/squirrel"
	"github.com/infinity-ocean/goldconv/internal/model"
)

func (r *repo) SelectBalance(id int) (model.Balance, error) {
	conn, err := r.pool.Acquire(context.Background()) // 1 select -> rows
	if err != nil {
		return model.Balance{}, fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()
	// TODO Добавить squirell (забумбенить сэнди)
	// builder := sq.Select("balance").From("coins") .Join("users USING (userFK)") 
	// builder.Where(sq.Eq{"userFK": id})
	// sql, args, err := builder.ToSql()
	// if err != nil {
	// 	return model.Balance{}, fmt.Errorf("failed to build SQL query: %w", err)
	// }
	sql := fmt.Sprintf("SELECT balance FROM coins WHERE userFK = %d", id)
	
	row := conn.QueryRow(context.Background(), sql) // args
	balanceScan := coins{} 
	err = row.Scan(&balanceScan.balance) // 2 rows -> balanceInt
	if err != nil {
		return model.Balance{}, fmt.Errorf("failed to scan result: %w", err)
	}
	var balance model.Balance // 3 balanceInt -> balance{g,s,b}
	balanceInt := balanceScan.balance
	balance.Gold = balanceInt / 10000
	balance.Silver = uint8((balanceInt % 10000) / 100) 
	balance.Bronze = uint8(balanceInt % 100)
	return balance, nil
}