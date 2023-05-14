package dto

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
