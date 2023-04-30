package repository

import (
	"Invalytics/app/pkg/model"
	"database/sql"
)

type Authorization interface {
	CreateUser(user model.User) (int32, error)
	GetUser(username, password string) (model.User, error)
}

type Deposit interface {
}

type Repository struct {
	Authorization
	Deposit
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
