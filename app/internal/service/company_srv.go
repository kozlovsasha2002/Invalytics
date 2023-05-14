package service

import (
	"Invalytics/app/internal/dto"
	"Invalytics/app/internal/model"
	"Invalytics/app/internal/repository"
	"sort"
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

func (s *CompanyService) GetAllMultipliers(userId int32, param string) ([]dto.MultiplierDto, error) {
	companies, err := s.repo.GetAllCompanies(userId)
	if err != nil {
		return nil, err
	}
	multipliers := make([]dto.MultiplierDto, 0, len(companies))
	for _, comp := range companies {
		mult := CalculateMultipliers(comp)
		multipliers = append(multipliers, mult)
	}

	sort.Slice(multipliers, func(i, j int) bool {
		if param == "ebitda" {
			return multipliers[i].EBITDA > multipliers[j].EBITDA
		} else if param == "pe" {
			return multipliers[i].PE < multipliers[j].PE
		} else {
			return multipliers[i].EV > multipliers[j].EV
		}
	})
	return multipliers, nil
}

func (s *CompanyService) GetMultiplierById(userId, compId int32) (dto.MultiplierDto, error) {
	var mult dto.MultiplierDto
	comp, err := s.repo.GetCompanyById(userId, compId)
	if err != nil {
		return mult, err
	}
	mult = CalculateMultipliers(comp)
	return mult, nil
}

func CalculateMultipliers(comp model.Company) dto.MultiplierDto {
	var mult dto.MultiplierDto
	mult.Name = comp.Name
	mult.EBITDA = CalculateEBITDA(comp)
	mult.PE = CalculatePE(comp)
	mult.EV = CalculateEV(comp)
	return mult
}

func CalculateEBITDA(comp model.Company) float32 {
	return float32(comp.Revenue - comp.TransactionCosts)
}

func CalculatePE(comp model.Company) float32 {
	if comp.AnnualProfit <= 0 {
		panic("annual profit < 0")
	}
	return float32(comp.MarketCapitalization) / float32(comp.AnnualProfit)
}

func CalculateEV(comp model.Company) float32 {
	return float32(comp.MarketCapitalization + comp.Debentures - comp.AvailableFunds)
}
