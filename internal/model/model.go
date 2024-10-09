package model

import "time"

type Balance struct {
	Gold   uint
	Silver uint8
	Bronze uint8
}

type Account struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Number    int64     `json:"number"`
	CreatedAt time.Time `json:"createdAt"`
	Balance   int64     `json:"balance"`
}

type AccountSmall struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Balance  string `json:"balance"`
}
