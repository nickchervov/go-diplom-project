package domain

import "github.com/golang-jwt/jwt/v5"

type Task struct {
	Id      string `db:"id" json:"id"`
	Date    string `db:"date" json:"date"`
	Title   string `db:"title" json:"title"`
	Comment string `db:"comment" json:"comment"`
	Repeat  string `db:"repeat" json:"repeat"`
}

type Claims struct {
	PasswordHash [32]byte
	jwt.RegisteredClaims
}

func NewTask(date, title, comment, repeat string) Task {
	return Task{
		Title:   title,
		Comment: comment,
		Repeat:  repeat,
		Date:    date,
	}
}
