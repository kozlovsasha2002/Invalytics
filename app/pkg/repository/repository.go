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
	Create(userId int32, deposit model.Deposit) (int32, error)
	GetAll(userId int32) ([]model.Deposit, error)
	GetById(userId, id int32) (model.Deposit, error)
	Update(userId, id int32, input model.UpdateDeposit) error
	Delete(userId, id int32) error
}

type Repository struct {
	Authorization
	Deposit
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Deposit:       NewDepositPostgres(db),
	}
}
