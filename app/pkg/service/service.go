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
	Create(userId int32, deposit model.Deposit) (int32, error)
	GetAll(userId int32) ([]model.Deposit, error)
	GetById(userId, id int32) (model.Deposit, error)
	Update(userId, id int32, deposit model.UpdateDeposit) error
	Delete(userId, id int32) error
}

type Service struct {
	Authorization
	Deposit
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Deposit:       NewDepositService(repo),
	}
}
