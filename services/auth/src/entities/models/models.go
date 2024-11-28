package models

type User struct {
	Id int
	Email string
	Password_Hash []byte
}