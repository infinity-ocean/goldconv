package repo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
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

func (r *repo) Login(acc model.AccountLogin) (string, int, error) {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return "", 0, fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()
	sql := fmt.Sprintln("SELECT id, username, password FROM accounts WHERE username = $1")
	accScan := accountLogin{}
	err = conn.QueryRow(context.Background(), sql, acc.Username).Scan(&accScan.id, &accScan.username, &accScan.password)
	if err != nil {
		return "", 0, fmt.Errorf("failed to scan result: %w", err)
	}
	if accScan.username != acc.Username && accScan.password != acc.Password {
		return "", 0, errors.New("username or password are invalid")
	}
	jwt, err := createJWT(accScan.id)
	return jwt, accScan.id, err
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

func (r *repo) SelectAccount(id int) (model.Account, error) {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return model.Account{}, fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()
	sql := fmt.Sprintln("SELECT * FROM accounts WHERE id = $1")
	accScan := account{}
	err = conn.QueryRow(context.Background(), sql, id).Scan(
		&accScan.ID,
		&accScan.Username,
		&accScan.Email,
		&accScan.Password,
		&accScan.Number,
		&accScan.CreatedAt,
		&accScan.Balance,
	)
	if err != nil {
		return model.Account{}, err
	}
	if hasEmptyFields(accScan){
		return model.Account{}, errors.New("retrieved account from PG has empty fields")
	}
	return accScan.toAccount(), nil
}

func hasEmptyFields(s interface{}) bool {
    val := reflect.ValueOf(s)

    if val.Kind() != reflect.Struct {
        return false
    }

    for i := 0; i < val.NumField(); i++ {
        if isEmpty(val.Field(i)) {
            return true // Найдено пустое поле
        }
    }
    return false // Пустых полей нет
}

// Вспомогательная функция для проверки на пустоту
func isEmpty(v reflect.Value) bool {
    switch v.Kind() {
    case reflect.String:
        return v.String() == ""
    case reflect.Int:
        return v.Int() == 0
    case reflect.Slice, reflect.Map, reflect.Array:
        return v.IsNil() || v.Len() == 0
    default:
        return false
    }
}