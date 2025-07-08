package models

type User struct {
	ID        int
	Login     string
	Password  string
	Salt      string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}
