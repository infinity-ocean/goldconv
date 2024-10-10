package repo

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/infinity-ocean/goldconv/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
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

func (r *repo) Login(acc model.AccountLogin) (string, error) {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()
	sql := fmt.Sprintln("SELECT id, username, password FROM accounts WHERE username = $1")
	accScan := accountLogin{}
	err = conn.QueryRow(context.Background(), sql, acc.Username).Scan(&accScan.id, &accScan.username, &accScan.password)
	if err != nil {
		return "", fmt.Errorf("failed to scan result: %w", err)
	}
	if accScan.username != acc.Username && accScan.password != acc.Password {
		return "", errors.New("username or password are invalid")
	}

	return createJWT(accScan.id)
}

func createJWT(id int) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"accountID": id,
	}
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
