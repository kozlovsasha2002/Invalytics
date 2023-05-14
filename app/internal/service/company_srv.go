package service

import (
	"Invalytics/app/internal/dto"
	"Invalytics/app/internal/model"
	"Invalytics/app/internal/repository"
)

type CompanyService struct {
	repo repository.Company
}

func NewCompanyService(repo repository.Company) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) CreateCompany(userId int32, comp model.Company) (int32, error) {
	return s.repo.CreateCompany(userId, comp)
}

func (s *CompanyService) GetAllCompanies(userId int32) ([]model.Company, error) {
	return s.repo.GetAllCompanies(userId)
}

func (s *CompanyService) GetCompanyById(userId, compId int32) (model.Company, error) {
	return s.repo.GetCompanyById(userId, compId)
}

func (s *CompanyService) UpdateCompany(userId, compId int32, input dto.UpdateCompanyDto) error {
	if err := input.Validate(); err != nil {
		return err
	}
	if _, err := s.repo.GetCompanyById(userId, compId); err != nil {
		return err
	}
	return s.repo.UpdateCompany(userId, compId, input)
}

func (s *CompanyService) DeleteCompany(userId, compId int32) error {
	if _, err := s.repo.GetCompanyById(userId, compId); err != nil {
		return err
	}
	return s.repo.DeleteCompany(userId, compId)
}

func (s *CompanyService) GetAllMultipliers(userId int32) ([]dto.MultiplierDto, error) {
	return nil, nil
}

func (s *CompanyService) GetMultiplierById(userId, compId int32) (dto.MultiplierDto, error) {
	return dto.MultiplierDto{}, nil
}
