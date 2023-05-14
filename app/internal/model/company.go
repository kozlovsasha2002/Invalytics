package model

type Company struct {
	Id                   int32  `json:"id"`
	Name                 string `json:"name"`
	DeptPayments         int    `json:"deptPayments"`         //выплаты по долгам (проценты)
	Depreciation         int    `json:"depreciation"`         //амортизация
	Taxes                int    `json:"taxes"`                //налоги
	MarketCapitalization int    `json:"marketCapitalization"` //рыночная капитализация
	AnnualProfit         int    `json:"annualProfit"`         //годовая прибыль
	Debentures           int    `json:"debentures"`           //все долговые обязательства
}

/*
Мультипликатор EBITDA — это прибыль компании до выплаты процентов, налогов и амортизации.
P/E — отношение цены компании к прибыли, а если точнее, рыночной капитализации всей компании к годовой чистой прибыли.
Мультипликатор EV — это справедливая стоимость компании. Определяется так:
EV = Рыночная капитализация + Все долговые обязательства − Доступные денежные средства компании.
*/
