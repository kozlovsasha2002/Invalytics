package repository

import (
	"Invalytics/app/internal/model"
	"Invalytics/app/pkg/postgresql"
	"database/sql"
	"fmt"
	"strings"
)

type BondPostgres struct {
	db *sql.DB
}

func NewBondPostgres(db *sql.DB) *BondPostgres {
	return &BondPostgres{db: db}
}

func (r *BondPostgres) CreateBond(userId int32, bond model.Bond) (int32, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int32
	bondQuery := fmt.Sprintf("INSERT INTO %s (ticker, amount_of_months, redemption_date, size_of_coupon, number_of_payments, purchase_price, nominal) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", postgresql.BondsTable)
	row := r.db.QueryRow(bondQuery, bond.Ticker, bond.AmountOfMonths, bond.RedemptionDate, bond.SizeOfCoupon, bond.NumberOfPayments, bond.PurchasePrice, bond.Nominal)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	userBondQuery := fmt.Sprintf("INSERT INTO %r (user_id, bond_id) VALUES ($1, $2)", postgresql.UserBondTable)
	_, err = tx.Exec(userBondQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BondPostgres) GetAllBonds(userId int32) ([]model.Bond, error) {
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

func (r *BondPostgres) GetBondById(userId, id int32) (model.Bond, error) {
	var bond model.Bond
	query := fmt.Sprintf("SELECT bond.id, bond.ticker, bond.amount_of_months, bond.redemption_date, bond.size_of_coupon, bond.number_of_payments, bond.purchase_price, bond.nominal FROM %s bond INNER JOIN %s userbonds ON bond.id = userbonds.bond_id WHERE userbonds.user_id = $1 AND userbonds.bond_id = $2", postgresql.BondsTable, postgresql.UserBondTable)

	row := r.db.QueryRow(query, userId, id)
	if err := row.Scan(&bond.Id, &bond.Ticker, &bond.AmountOfMonths, &bond.RedemptionDate, &bond.SizeOfCoupon, &bond.NumberOfPayments, &bond.PurchasePrice, &bond.Nominal); err != nil {
		return bond, err
	}

	return bond, nil
}

func (r *BondPostgres) UpdateBond(userId, id int32, input model.UpdateBond) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Ticker != nil {
		setValues = append(setValues, fmt.Sprintf("ticker=$%d", argId))
		args = append(args, *input.Ticker)
		argId++
	}
	if input.AmountOfMonths != nil {
		setValues = append(setValues, fmt.Sprintf("amount_of_months=$%d", argId))
		args = append(args, *input.AmountOfMonths)
		argId++
	}
	if input.RedemptionDate != nil {
		setValues = append(setValues, fmt.Sprintf("redemption_date=$%d", argId))
		args = append(args, *input.RedemptionDate)
		argId++
	}
	if input.SizeOfCoupon != nil {
		setValues = append(setValues, fmt.Sprintf("size_of_coupon=$%d", argId))
		args = append(args, *input.SizeOfCoupon)
		argId++
	}
	if input.NumberOfPayments != nil {
		setValues = append(setValues, fmt.Sprintf("number_of_payments=$%d", argId))
		args = append(args, *input.NumberOfPayments)
		argId++
	}
	if input.PurchasePrice != nil {
		setValues = append(setValues, fmt.Sprintf("purchase_price=$%d", argId))
		args = append(args, *input.PurchasePrice)
		argId++
	}
	if input.Nominal != nil {
		setValues = append(setValues, fmt.Sprintf("nominal=$%d", argId))
		args = append(args, *input.Nominal)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s bonds SET %s FROM %s userbonds WHERE bonds.id = userbonds.bond_id AND userbonds.user_id = $%d AND userbonds.bond_id = $%d",
		postgresql.BondsTable, setQuery, postgresql.UserBondTable, argId, argId+1)
	args = append(args, userId, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *BondPostgres) DeleteBond(userId, id int32) error {
	query := fmt.Sprintf("DELETE FROM %s bonds USING %s userbonds WHERE bonds.id = userbonds.bond_id AND userbonds.user_id = $1 AND userbonds.bond_id = $2", postgresql.BondsTable, postgresql.UserBondTable)
	_, err := r.db.Exec(query, userId, id)
	return err
}
