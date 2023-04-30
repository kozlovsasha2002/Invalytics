package service

import (
	"Invalytics/app/pkg/model"
	"Invalytics/app/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int32, error)
	GenerateToken(username, password string) (string, error)
}

type Deposit interface {
}

type Service struct {
	Authorization
	Deposit
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
