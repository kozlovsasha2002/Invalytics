package service

import (
	"Invalytics/app/pkg/model"
	"Invalytics/app/pkg/repository"
)

type DepositService struct {
	repo repository.Deposit
}

func NewDepositService(repo repository.Deposit) *DepositService {
	return &DepositService{repo: repo}
}

func (s *DepositService) Create(userId int32, deposit model.Deposit) (int32, error) {
	return s.repo.Create(userId, deposit)
}

func (s *DepositService) GetAll(userId int32) ([]model.Deposit, error) {
	return s.repo.GetAll(userId)
}

func (s *DepositService) GetById(userId, id int32) (model.Deposit, error) {
	return s.repo.GetById(userId, id)
}

func (s *DepositService) Update(userId, id int32, deposit model.UpdateDeposit) error {
	if err := deposit.Validate(); err != nil {
		return err
	}

	if _, err := s.repo.GetById(userId, id); err != nil {
		return err
	}

	return s.repo.Update(userId, id, deposit)
}

func (s *DepositService) Delete(userId, id int32) error {
	if _, err := s.repo.GetById(userId, id); err != nil {
		return err
	}

	return s.repo.Delete(userId, id)
}
