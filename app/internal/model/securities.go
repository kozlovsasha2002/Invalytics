package model

type Securities interface {
	ProfitabilityById(term int) (float32, error)
}

type ProfitInfo struct {
	Ticker         string
	Profit         float32
	AmountOfMonths int
}
