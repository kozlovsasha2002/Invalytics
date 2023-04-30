package model

import "time"

type Deposit struct {
	Id             int32     `json:"id"`
	InitialAmount  int       `json:"initialAmount"`
	StartDate      time.Time `json:"startDate"`
	NumberOfMonths int       `json:"numberOfMonths"`
	PercentageRate float32   `json:"percentageRate"`
}

func (d *Deposit) EndDate() time.Time {

	return time.Time{}
}

func (d *Deposit) FinalAmount() int {
	// итоговая сумма после погашения депозита
	return 0
}

func (d *Deposit) AnnualReturn() float32 {
	// доходность в процентах годовых
	return d.PercentageRate
}

func (d *Deposit) CapitalGain() float32 {
	//разница между вложенной суммой и полученной
	return 0
}

func (d *Deposit) CapitalGainInPercent() float32 {
	// отношения полученной суммы к вложенной минус 100 %
	return 0
}
