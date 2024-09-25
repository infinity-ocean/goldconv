package repo

import (
	"context"
	"fmt"

	// sq "github.com/Masterminds/squirrel"
	"github.com/infinity-ocean/goldconv/internal/model"
)

func (r *repo) SelectBalance(id int) (model.Balance, error) {
	conn, err := r.pool.Acquire(context.Background()) // 1
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
	err = row.Scan(&balanceScan.balance) // 2
	if err != nil {
		return model.Balance{}, fmt.Errorf("failed to scan result: %w", err)
	}
	balanceInt := balanceScan.balance
	var balance model.Balance // 3
	// TODO система работает неверно
	balance.Gold = balanceInt / 100
	balance.Silver = uint8((balanceInt % 100) / 10) * 10
	balance.Bronze = uint8(balanceInt % 100)
	return balance, nil
}