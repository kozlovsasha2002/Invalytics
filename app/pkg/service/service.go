package service

import (
	"Invalytics/app/pkg/model"
	"Invalytics/app/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int32, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int32, error)
}

type Deposit interface {
	CreateDeposit(userId int32, deposit model.Deposit) (int32, error)
	GetAllDeposits(userId int32) ([]model.Deposit, error)
	GetDepositById(userId, id int32) (model.Deposit, error)
	UpdateDeposit(userId, id int32, deposit model.UpdateDeposit) error
	DeleteDeposit(userId, id int32) error
}

type Bond interface {
	CreateBond(userId int32, bond model.Bond) (int32, error)
	GetAllBonds(userId int32) ([]model.Bond, error)
	GetBondById(userId, id int32) (model.Bond, error)
	UpdateBond(userId, id int32, input model.UpdateBond) error
	DeleteBond(userId, id int32) error
}

type Service struct {
	Authorization
	Deposit
	Bond
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Deposit:       NewDepositService(repo),
		Bond:          NewBondService(repo),
	}
}
