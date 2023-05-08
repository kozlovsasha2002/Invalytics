package service

import (
	"Invalytics/app/internal/model"
	"Invalytics/app/internal/repository"
)

type ShareService struct {
	repo repository.Share
}

func NewShareService(repo repository.Share) *ShareService {
	return &ShareService{repo: repo}
}

func (s *ShareService) CreateShare(userId int32, input model.Share) (int32, error) {
	return s.repo.CreateShare(userId, input)
}

func (s *ShareService) GetAllShares(userId int32) ([]model.Share, error) {
	return s.repo.GetAllShares(userId)
}

func (s *ShareService) GetShareById(userId, id int32) (model.Share, error) {
	return s.repo.GetShareById(userId, id)
}

func (s *ShareService) UpdateShare(userId, id int32, input model.UpdateShare) error {
	if err := input.Validate(); err != nil {
		return err
	}

	if _, err := s.repo.GetShareById(userId, id); err != nil {
		return err
	}

	return s.repo.UpdateShare(userId, id, input)
}

func (s *ShareService) DeleteShare(userId, id int32) error {
	if _, err := s.repo.GetShareById(userId, id); err != nil {
		return err
	}

	return s.repo.DeleteShare(userId, id)
}
