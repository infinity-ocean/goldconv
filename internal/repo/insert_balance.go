package repo

import (
	"context"
	"fmt"

	"github.com/infinity-ocean/goldconv/internal/model"
)

func (r *repo) InsertBalance(id int, balance model.Balance) error{
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()
	var exists bool
	err = conn.QueryRow(context.Background(),"SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", id).Scan(&exists)
	if err != nil {
	    return fmt.Errorf("failed select user: %w", err)
	}
	var balanceUint uint
	balanceUint = balance.Gold * 10000 // 54350000
	balanceUint += uint(balance.Silver) * 100 // 54350000 + 7800 = 54357800
	balanceUint += uint(balance.Bronze) // 54357800 + 29 = 54357829
	if exists {
    	sql := fmt.Sprintf(
    	    `INSERT INTO coins (userFK, balance) 
			VALUES (%d, %d) 
			ON CONFLICT (userFK) 
			DO UPDATE SET balance = EXCLUDED.balance;`,
    	    id, balanceUint,
		)
    	_, err = conn.Exec(context.Background(), sql)
    	if err != nil {
    	    return fmt.Errorf("can't Upsert coins: %w", err)

    	}
	} else {
		return fmt.Errorf("user is not exists")
	}
	return nil
}