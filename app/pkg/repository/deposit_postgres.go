package repository

import (
	"Invalytics/app/pkg/model"
	"database/sql"
	"fmt"
	"strings"
)

type DepositPostgres struct {
	db *sql.DB
}

func NewDepositPostgres(db *sql.DB) *DepositPostgres {
	return &DepositPostgres{db: db}
}

func (r *DepositPostgres) CreateDeposit(userId int32, deposit model.Deposit) (int32, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int32
	depositQuery := fmt.Sprintf("INSERT INTO %s (initial_amount, start_date, number_of_months, percentage_rate) VALUES ($1, $2, $3, $4) RETURNING id", depositsTable)
	row := tx.QueryRow(depositQuery, deposit.InitialAmount, deposit.StartDate, deposit.NumberOfMonths, deposit.PercentageRate)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	userDepositQuery := fmt.Sprintf("INSERT INTO %s (user_id, deposit_id) VALUES ($1, $2)", userDepositTable)
	_, err = tx.Exec(userDepositQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *DepositPostgres) GetAllDeposits(userId int32) ([]model.Deposit, error) {
	list := make([]model.Deposit, 0)
	query := fmt.Sprintf("SELECT deps.id, deps.initial_amount, deps.start_date, deps.number_of_months, deps.percentage_rate FROM %s deps INNER JOIN %s userdeps ON deps.id = userdeps.deposit_id WHERE userdeps.user_id = $1",
		depositsTable, userDepositTable)
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

func (r *DepositPostgres) GetDepositById(userId, id int32) (model.Deposit, error) {
	var d model.Deposit
	query := fmt.Sprintf("SELECT deps.id, deps.initial_amount, deps.start_date, deps.number_of_months, deps.percentage_rate FROM %s deps INNER JOIN %s userdeps ON deps.id = userdeps.deposit_id WHERE userdeps.user_id = $1 AND userdeps.deposit_id = $2",
		depositsTable, userDepositTable)

	row := r.db.QueryRow(query, userId, id)
	if err := row.Scan(&d.Id, &d.InitialAmount, &d.StartDate, &d.NumberOfMonths, &d.PercentageRate); err != nil {
		return d, err
	}

	return d, nil
}

func (r *DepositPostgres) UpdateDeposit(userId, id int32, input model.UpdateDeposit) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.InitialAmount != nil {
		setValues = append(setValues, fmt.Sprintf("initial_amount=$%d", argId))
		args = append(args, *input.InitialAmount)
		argId++
	}

	if input.StartDate != nil {
		setValues = append(setValues, fmt.Sprintf("start_date=$%d", argId))
		args = append(args, *input.StartDate)
		argId++
	}

	if input.NumberOfMonths != nil {
		setValues = append(setValues, fmt.Sprintf("number_of_months=$%d", argId))
		args = append(args, *input.NumberOfMonths)
		argId++
	}

	if input.PercentageRate != nil {
		setValues = append(setValues, fmt.Sprintf("percentage_rate=$%d", argId))
		args = append(args, *input.PercentageRate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s deps SET %s FROM %s userdeps WHERE deps.id = userdeps.deposit_id AND userdeps.user_id=$%d AND userdeps.deposit_id=$%d",
		depositsTable, setQuery, userDepositTable, argId, argId+1)
	args = append(args, userId, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *DepositPostgres) DeleteDeposit(userId, id int32) error {
	query := fmt.Sprintf("DELETE FROM %s deps USING %s userdeps WHERE deps.id = userdeps.deposit_id AND userdeps.user_id = $1 AND userdeps.deposit_id = $2",
		depositsTable, userDepositTable)
	_, err := r.db.Exec(query, userId, id)
	return err
}
