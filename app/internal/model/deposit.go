package model

import (
	"errors"
)

type Deposit struct {
	Id             int32   `json:"id"`
	InitialAmount  int     `json:"initialAmount"`
	StartDate      string  `json:"startDate"`
	NumberOfMonths int     `json:"numberOfMonths"`
	PercentageRate float32 `json:"percentageRate"`
}

func (d *Deposit) ProfitabilityById(term int) (float32, error) {
	var profit float32
	if term <= 0 {
		return 0, errors.New("term <= 0")
	}
	sum := float32(d.InitialAmount) * d.PercentageRate
	years := float32(term) / float32(12)
	profit = (sum * years) / float32(d.InitialAmount)
	return profit, nil
}

type UpdateDeposit struct {
	InitialAmount  *int     `json:"initialAmount"`
	StartDate      *string  `json:"startDate"`
	NumberOfMonths *int     `json:"numberOfMonths"`
	PercentageRate *float32 `json:"percentageRate"`
}

func (u *UpdateDeposit) Validate() error {
	if u.InitialAmount == nil && u.StartDate == nil && u.NumberOfMonths == nil && u.PercentageRate == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
