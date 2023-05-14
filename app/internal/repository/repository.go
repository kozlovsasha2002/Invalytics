package repository

import (
	"Invalytics/app/internal/model"
	"database/sql"
)

type Repository struct {
	Authorization
	Deposit
	Bond
	Share
	Profit
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Deposit:       NewDepositPostgres(db),
		Bond:          NewBondPostgres(db),
		Share:         NewSharePostgres(db),
		Profit:        NewProfitPostgres(db),
	}
}

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

type Share interface {
	CreateShare(userId int32, share model.Share) (int32, error)
	GetAllShares(userId int32) ([]model.Share, error)
	GetShareById(userId, id int32) (model.Share, error)
	UpdateShare(userId, id int32, input model.UpdateShare) error
	DeleteShare(userId, id int32) error
}

type Profit interface {
	GetShare(userId, id int32) (model.Share, error)
	AllShares(userId int32) ([]model.Share, error)
	GetBond(userId, id int32) (model.Bond, error)
	AllBonds(userId int32) ([]model.Bond, error)
	GetDeposit(userId, id int32) (model.Deposit, error)
	AllDeposits(userId int32) ([]model.Deposit, error)
}
