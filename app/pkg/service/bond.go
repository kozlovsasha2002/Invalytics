package service

import (
	"Invalytics/app/pkg/model"
	"Invalytics/app/pkg/repository"
)

type BondService struct {
	repo repository.Bond
}

func NewBondService(repo repository.Bond) *BondService {
	return &BondService{repo: repo}
}

func (s *BondService) CreateBond(userId int32, bond model.Bond) (int32, error) {
	return s.repo.CreateBond(userId, bond)
}

func (s *BondService) GetAllBonds(userId int32) ([]model.Bond, error) {
	return s.repo.GetAllBonds(userId)
}

func (s *BondService) GetBondById(userId, id int32) (model.Bond, error) {
	return s.repo.GetBondById(userId, id)
}

// UpdateBond todo
func (s *BondService) UpdateBond(userId, id int32, input model.UpdateBond) error {
	if err := input.Validate(); err != nil {
		return err
	}

	if _, err := s.repo.GetBondById(userId, id); err != nil {
		return err
	}

	return s.repo.UpdateBond(userId, id, input)
}

// DeleteBond todo
func (s *BondService) DeleteBond(userId, id int32) error {
	if _, err := s.repo.GetBondById(userId, id); err != nil {
		return err
	}
	return s.repo.DeleteBond(userId, id)
}
