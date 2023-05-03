package model

import "errors"

type Bond struct {
	Id               int32   `json:"id"`
	Ticker           string  `json:"ticker"`
	AmountOfMonths   int     `json:"amountOfMonths"`
	RedemptionDate   string  `json:"redemptionDate"`
	SizeOfCoupon     float32 `json:"sizeOfCoupon"`
	NumberOfPayments int8    `json:"numberOfPayments"`
	PurchasePrice    float32 `json:"purchasePrice"`
	Nominal          int16   `json:"nominal"`
}

type UpdateBond struct {
	Ticker           *string  `json:"ticker"`
	AmountOfMonths   *int     `json:"amountOfMonths"`
	RedemptionDate   *string  `json:"redemptionDate"`
	SizeOfCoupon     *float32 `json:"sizeOfCoupon"`
	NumberOfPayments *int8    `json:"numberOfPayments"`
	PurchasePrice    *float32 `json:"purchasePrice"`
	Nominal          *int16   `json:"nominal"`
}

func (u *UpdateBond) Validate() error {
	if u.Ticker == nil && u.AmountOfMonths == nil && u.RedemptionDate == nil && u.SizeOfCoupon == nil && u.NumberOfPayments == nil && u.PurchasePrice == nil && u.Nominal == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
