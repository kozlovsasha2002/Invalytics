package dto

import "errors"

type BondInfo struct {
	Ticker         string
	Profit         float32
	AmountOfMonths int
}

type DepositDto struct {
	PercentageRate float32
	Profit         float32
	AmountOfMonths int
}

type MultiplierDto struct {
	Name   string
	PE     float32
	EBITDA float32
	EV     float32
}

type UpdateCompanyDto struct {
	Name                 *string `json:"name"`
	DeptPayments         *int    `json:"deptPayments"`
	Depreciation         *int    `json:"depreciation"`
	Taxes                *int    `json:"taxes"`
	MarketCapitalization *int    `json:"marketCapitalization"`
	AnnualProfit         *int    `json:"annualProfit"`
	Debentures           *int    `json:"debentures"`
	Revenue              *int    `json:"revenue"`
	TransactionCosts     *int    `json:"transactionCosts"`
	AvailableFunds       *int    `json:"availableFunds"`
}

func (u *UpdateCompanyDto) Validate() error {
	if u.Name == nil && u.DeptPayments == nil && u.Depreciation == nil && u.MarketCapitalization == nil && u.AnnualProfit == nil && u.Debentures == nil &&
		u.Revenue == nil && u.TransactionCosts == nil && u.AvailableFunds == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
