package repository

import (
	"Invalytics/app/internal/model"
	"Invalytics/app/pkg/postgresql"
	"database/sql"
	"fmt"
)

type ProfitPostgres struct {
	db *sql.DB
}

func NewProfitPostgres(db *sql.DB) *ProfitPostgres {
	return &ProfitPostgres{db: db}
}

func (r *ProfitPostgres) GetShare(userId, id int32) (model.Share, error) {
	var share model.Share
	query := fmt.Sprintf("SELECT s.id, s.ticker, s.purchase_price, s.estimated_selling_price, s.expected_amount_of_dividends, s.amount_of_months FROM %s s INNER JOIN %s u ON s.id = u.share_id WHERE u.user_id = $1 AND u.share_id = $2", postgresql.SharesTable, postgresql.UserShareTable)

	row := r.db.QueryRow(query, userId, id)
	if err := row.Scan(&share.Id, &share.Ticker, &share.PurchasePrice, &share.EstimatedSellingPrice, &share.ExpectedAmountOfDividends, &share.AmountOfMonths); err != nil {
		return share, err
	}

	return share, nil
}

func (r *ProfitPostgres) AllShares(userId int32) ([]model.Share, error) {
	shares := make([]model.Share, 0)
	query := fmt.Sprintf("SELECT s.id, s.ticker, s.purchase_price, s.estimated_selling_price, s.expected_amount_of_dividends, s.amount_of_months FROM %s s INNER JOIN %s u ON s.id = u.share_id WHERE u.user_id = $1", postgresql.SharesTable, postgresql.UserShareTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s model.Share
		err := rows.Scan(&s.Id, &s.Ticker, &s.PurchasePrice, &s.EstimatedSellingPrice, &s.ExpectedAmountOfDividends, &s.AmountOfMonths)
		if err != nil {
			return nil, err
		}
		shares = append(shares, s)
	}

	return shares, nil
}

func (r *ProfitPostgres) GetBond(userId, id int32) (model.Bond, error) {
	var bond model.Bond
	query := fmt.Sprintf("SELECT bond.id, bond.ticker, bond.amount_of_months, bond.redemption_date, bond.size_of_coupon, bond.number_of_payments, bond.purchase_price, bond.nominal FROM %s bond INNER JOIN %s userbonds ON bond.id = userbonds.bond_id WHERE userbonds.user_id = $1 AND userbonds.bond_id = $2", postgresql.BondsTable, postgresql.UserBondTable)

	row := r.db.QueryRow(query, userId, id)
	if err := row.Scan(&bond.Id, &bond.Ticker, &bond.AmountOfMonths, &bond.RedemptionDate, &bond.SizeOfCoupon, &bond.NumberOfPayments, &bond.PurchasePrice, &bond.Nominal); err != nil {
		return bond, err
	}

	return bond, nil
}

func (r *ProfitPostgres) AllBonds(userId int32) ([]model.Bond, error) {
	bonds := make([]model.Bond, 0)
	query := fmt.Sprintf("SELECT b.id, b.ticker, b.amount_of_months, b.redemption_date, b.size_of_coupon, b.number_of_payments, b.purchase_price, b.nominal FROM %s b INNER JOIN %s userbonds ON b.id = userbonds.bond_id WHERE userbonds.user_id = $1", postgresql.BondsTable, postgresql.UserBondTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bond model.Bond
		err := rows.Scan(&bond.Id, &bond.Ticker, &bond.AmountOfMonths, &bond.RedemptionDate, &bond.SizeOfCoupon, &bond.NumberOfPayments, &bond.PurchasePrice, &bond.Nominal)
		if err != nil {
			return nil, err
		}
		bonds = append(bonds, bond)
	}

	return bonds, nil
}

func (r *ProfitPostgres) GetDeposit(userId, id int32) (model.Deposit, error) {
	var d model.Deposit
	query := fmt.Sprintf("SELECT deps.id, deps.initial_amount, deps.start_date, deps.number_of_months, deps.percentage_rate FROM %s deps INNER JOIN %s userdeps ON deps.id = userdeps.deposit_id WHERE userdeps.user_id = $1 AND userdeps.deposit_id = $2",
		postgresql.DepositsTable, postgresql.UserDepositTable)

	row := r.db.QueryRow(query, userId, id)
	if err := row.Scan(&d.Id, &d.InitialAmount, &d.StartDate, &d.NumberOfMonths, &d.PercentageRate); err != nil {
		return d, err
	}

	return d, nil
}

func (r *ProfitPostgres) AllDeposits(userId int32) ([]model.Deposit, error) {
	list := make([]model.Deposit, 0)
	query := fmt.Sprintf("SELECT deps.id, deps.initial_amount, deps.start_date, deps.number_of_months, deps.percentage_rate FROM %s deps INNER JOIN %s userdeps ON deps.id = userdeps.deposit_id WHERE userdeps.user_id = $1",
		postgresql.DepositsTable, postgresql.UserDepositTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d model.Deposit
		err := rows.Scan(&d.Id, &d.InitialAmount, &d.StartDate, &d.NumberOfMonths, &d.PercentageRate)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}
