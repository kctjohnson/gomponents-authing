package models

type Account struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
