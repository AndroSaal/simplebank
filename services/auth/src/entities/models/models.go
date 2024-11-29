package models

type User struct {
	Id           int    `db:"id"`
	Email        string `db:"email"`
	PasswordHash []byte `db:"password_hash"`
}
