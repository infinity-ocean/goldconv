package repo

import (
	"time"

	"github.com/infinity-ocean/goldconv/internal/model"
)

type coins struct {
	balance uint `db:"balance"`
}

type accountLogin struct {
	id       int    `db:"id"`
	username string `db:"username"`
	password string `db:"password"`
}

type account struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Number    int     `db:"number"`
	CreatedAt time.Time `db:"createdAt"`
	Balance   int     `db:"balance"`
}

func (a account) toAccount() model.Account {
	return model.Account{
		ID:        a.ID,
		Username:  a.Username,
		Email:     a.Email,
		Password:  a.Password,
		Number:    a.Number,
		CreatedAt: a.CreatedAt,
		Balance:   a.Balance,
	}
}
