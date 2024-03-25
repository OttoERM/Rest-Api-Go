package models

type User struct {
	id         int
	username   string
	password   string
	created_at string
}

type Book struct {
	ID          int
	Title       string
	Description string
}
