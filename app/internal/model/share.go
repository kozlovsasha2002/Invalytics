package model

import "errors"

type Share struct {
	Id                        int32   `json:"id" db:"id"'`
	Ticker                    string  `json:"ticker"`
	PurchasePrice             float32 `json:"purchasePrice"`
	EstimatedSellingPrice     float32 `json:"estimatedSellingPrice"`
	ExpectedAmountOfDividends float32 `json:"expectedAmountOfDividends"`
	AmountOfMonths            int     `json:"amountOfMonths"`
}

func (s *Share) ProfitabilityById(termInMonths int) (float32, error) {
	if termInMonths < 0 {
		return 0, errors.New("amount of months less 0")
	}
	years := float32(termInMonths) / 12
	income := (s.EstimatedSellingPrice - s.PurchasePrice) + s.ExpectedAmountOfDividends
	percentage := (income / s.PurchasePrice) * 100 / years
	return percentage, nil
}

type UpdateShare struct {
	Ticker                    *string  `json:"ticker"`
	PurchasePrice             *float32 `json:"purchasePrice"`
	EstimatedSellingPrice     *float32 `json:"estimatedSellingPrice"`
	ExpectedAmountOfDividends *float32 `json:"expectedAmountOfDividends"`
	AmountOfMonths            *int     `json:"amountOfMonths"`
}

func (u *UpdateShare) Validate() error {
	if u.Ticker == nil && u.PurchasePrice == nil && u.EstimatedSellingPrice == nil && u.ExpectedAmountOfDividends == nil && u.AmountOfMonths == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
