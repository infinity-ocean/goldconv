package repo

type coins struct {
	balance uint `db:"balance"`
}

type accountLogin struct {
	id       int    `db:"id"`
	username string `db:"username"`
	password string `db:"password"`
}
