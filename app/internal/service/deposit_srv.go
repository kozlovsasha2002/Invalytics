package service

import (
	"Invalytics/app/internal/model"
	"Invalytics/app/internal/repository"
)

type DepositService struct {
	repo repository.Deposit
}

func NewDepositService(repo repository.Deposit) *DepositService {
	return &DepositService{repo: repo}
}

func (s *DepositService) CreateDeposit(userId int32, deposit model.Deposit) (int32, error) {
	return s.repo.CreateDeposit(userId, deposit)
}

func (s *DepositService) GetAllDeposits(userId int32) ([]model.Deposit, error) {
	return s.repo.GetAllDeposits(userId)
}

func (s *DepositService) GetDepositById(userId, id int32) (model.Deposit, error) {
	return s.repo.GetDepositById(userId, id)
}

func (s *DepositService) UpdateDeposit(userId, id int32, deposit model.UpdateDeposit) error {
	if err := deposit.Validate(); err != nil {
		return err
	}

	if _, err := s.repo.GetDepositById(userId, id); err != nil {
		return err
	}

	return s.repo.UpdateDeposit(userId, id, deposit)
}

func (s *DepositService) DeleteDeposit(userId, id int32) error {
	if _, err := s.repo.GetDepositById(userId, id); err != nil {
		return err
	}

	return s.repo.DeleteDeposit(userId, id)
}
