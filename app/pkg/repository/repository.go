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
	CreateDeposit(userId int32, deposit model.Deposit) (int32, error)
	GetAllDeposits(userId int32) ([]model.Deposit, error)
	GetDepositById(userId, id int32) (model.Deposit, error)
	UpdateDeposit(userId, id int32, input model.UpdateDeposit) error
	DeleteDeposit(userId, id int32) error
}

type Bond interface {
	CreateBond(userId int32, bond model.Bond) (int32, error)
	GetAllBonds(userId int32) ([]model.Bond, error)
	GetBondById(userId, id int32) (model.Bond, error)
	UpdateBond(userId, id int32, input model.UpdateBond) error
	DeleteBond(userId, id int32) error
}

type Repository struct {
	Authorization
	Deposit
	Bond
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Deposit:       NewDepositPostgres(db),
		Bond:          NewBondPostgres(db),
	}
}
