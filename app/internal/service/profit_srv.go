package service

import (
	"Invalytics/app/internal/dto"
	"Invalytics/app/internal/model"
	"Invalytics/app/internal/repository"
	"sort"
)

type ProfitService struct {
	repo repository.Profit
}

func NewProfitService(repo repository.Profit) *ProfitService {
	return &ProfitService{repo: repo}
}

func (s *ProfitService) ShareProfitabilityById(userId, shareId int32, termInMonths int) (model.ProfitInfo, error) {
	var result model.ProfitInfo
	share, err := s.repo.GetShare(userId, shareId)
	if err != nil {
		return result, err
	}
	result.Profit, err = share.ProfitabilityById(termInMonths)
	if err != nil {
		return result, err
	}
	result.Ticker = share.Ticker
	result.AmountOfMonths = termInMonths
	return result, nil
}

func (s *ProfitService) AllShareProfitability(userId int32, termInMonths int, srt bool) ([]model.ProfitInfo, error) {
	shares, err := s.repo.AllShares(userId)
	if err != nil {
		return nil, err
	}

	profits := make([]model.ProfitInfo, 0, len(shares))
	for _, s := range shares {
		prof, err := s.ProfitabilityById(termInMonths)
		if err != nil {
			return nil, err
		}
		info := model.ProfitInfo{Ticker: s.Ticker, Profit: prof, AmountOfMonths: termInMonths}
		profits = append(profits, info)
	}

	if srt == true {
		sort.Slice(profits, func(i, j int) bool {
			return profits[i].Profit > profits[j].Profit
		})
	}

	return profits, nil
}

func (s *ProfitService) BondProfitabilityById(userId, bondId int32) (dto.BondInfo, error) {
	var result dto.BondInfo
	bond, err := s.repo.GetBond(userId, bondId)
	if err != nil {
		return result, err
	}
	result.Profit, err = bond.ProfitabilityById(bond.AmountOfMonths)
	if err != nil {
		return result, err
	}
	result.Ticker = bond.Ticker
	result.AmountOfMonths = bond.AmountOfMonths
	return result, nil
}

func (s *ProfitService) AllBondProfitability(userId int32, srt bool) ([]dto.BondInfo, error) {
	bonds, err := s.repo.AllBonds(userId)
	if err != nil {
		return nil, err
	}

	profits := make([]dto.BondInfo, 0, len(bonds))
	for _, b := range bonds {
		prof, err := b.ProfitabilityById(b.AmountOfMonths)
		if err != nil {
			return nil, err
		}
		info := dto.BondInfo{Ticker: b.Ticker, Profit: prof, AmountOfMonths: b.AmountOfMonths}
		profits = append(profits, info)
	}

	if srt == true {
		sort.Slice(profits, func(i, j int) bool {
			return profits[i].Profit > profits[j].Profit
		})
	}

	return profits, nil
}

func (s *ProfitService) DepositProfitabilityById(userId, depositId int32) (dto.DepositDto, error) {
	var result dto.DepositDto
	deposit, err := s.repo.GetDeposit(userId, depositId)
	if err != nil {
		return result, err
	}

	result.Profit, err = deposit.ProfitabilityById(deposit.NumberOfMonths)
	if err != nil {
		return result, err
	}
	result.PercentageRate = deposit.PercentageRate
	result.AmountOfMonths = deposit.NumberOfMonths
	return result, nil
}

func (s *ProfitService) AllDepositProfitability(userId int32, srt bool) ([]dto.DepositDto, error) {
	deposits, err := s.repo.AllDeposits(userId)
	if err != nil {
		return nil, err
	}

	profits := make([]dto.DepositDto, 0, len(deposits))
	for _, b := range deposits {
		prof, err := b.ProfitabilityById(b.NumberOfMonths)
		if err != nil {
			return nil, err
		}
		info := dto.DepositDto{PercentageRate: b.PercentageRate, Profit: prof, AmountOfMonths: b.NumberOfMonths}
		profits = append(profits, info)
	}

	if srt == true {
		sort.Slice(profits, func(i, j int) bool {
			return profits[i].Profit > profits[j].Profit
		})
	}

	return profits, nil
}
