package repository

import (
	"Invalytics/app/pkg/model"
	"database/sql"
	"fmt"
	"strings"
)

type SharePostgres struct {
	db *sql.DB
}

func NewSharePostgres(db *sql.DB) *SharePostgres {
	return &SharePostgres{db: db}
}

func (r *SharePostgres) CreateShare(userId int32, share model.Share) (int32, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int32
	shareQuery := fmt.Sprintf("INSERT INTO %s (ticker, purchase_price, estimated_selling_price, expected_amount_of_dividends, amount_of_months) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		sharesTable)
	row := tx.QueryRow(shareQuery, share.Ticker, share.PurchasePrice, share.EstimatedSellingPrice, share.ExpectedAmountOfDividends, share.AmountOfMonths)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}

	userShareQuery := fmt.Sprintf("INSERT INTO %s (user_id, share_id) VALUES ($1, $2)", userShareTable)
	_, err = tx.Exec(userShareQuery, userId, id)
	if err != nil {
		return 0, nil
	}

	return id, tx.Commit()
}

func (r *SharePostgres) GetAllShares(userId int32) ([]model.Share, error) {
	shares := make([]model.Share, 0)
	query := fmt.Sprintf("SELECT s.id, s.ticker, s.purchase_price, s.estimated_selling_price, s.expected_amount_of_dividends, s.amount_of_months FROM %s s INNER JOIN %s u ON s.id = u.share_id WHERE u.user_id = $1", sharesTable, userShareTable)

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

func (r *SharePostgres) GetShareById(userId, id int32) (model.Share, error) {
	var share model.Share
	query := fmt.Sprintf("SELECT s.id, s.ticker, s.purchase_price, s.estimated_selling_price, s.expected_amount_of_dividends, s.amount_of_months FROM %s s INNER JOIN %s u ON s.id = u.share_id WHERE u.user_id = $1 AND u.share_id = $2", sharesTable, userShareTable)

	row := r.db.QueryRow(query, userId, id)
	if err := row.Scan(&share.Id, &share.Ticker, &share.PurchasePrice, &share.EstimatedSellingPrice, &share.ExpectedAmountOfDividends, &share.AmountOfMonths); err != nil {
		return share, err
	}

	return share, nil
}

func (r *SharePostgres) UpdateShare(userId, id int32, input model.UpdateShare) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Ticker != nil {
		setValues = append(setValues, fmt.Sprintf("ticker=$%d", argId))
		args = append(args, *input.Ticker)
		argId++
	}
	if input.PurchasePrice != nil {
		setValues = append(setValues, fmt.Sprintf("purchase_price=$%d", argId))
		args = append(args, *input.PurchasePrice)
		argId++
	}
	if input.EstimatedSellingPrice != nil {
		setValues = append(setValues, fmt.Sprintf("estimated_selling_price=$%d", argId))
		args = append(args, *input.EstimatedSellingPrice)
		argId++
	}
	if input.ExpectedAmountOfDividends != nil {
		setValues = append(setValues, fmt.Sprintf("expected_amount_of_dividends=$%d", argId))
		args = append(args, *input.ExpectedAmountOfDividends)
		argId++
	}
	if input.AmountOfMonths != nil {
		setValues = append(setValues, fmt.Sprintf("amount_of_months=$%d", argId))
		args = append(args, *input.AmountOfMonths)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s s SET %s FROM %s u WHERE s.id = u.share_id AND u.user_id = $%d AND u.share_id = $%d",
		sharesTable, setQuery, userShareTable, argId, argId+1)
	args = append(args, userId, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *SharePostgres) DeleteShare(userId, id int32) error {
	query := fmt.Sprintf("DELETE FROM %s s USING %s u WHERE s.id = u.share_id AND u.user_id = $1 AND u.share_id = $2", sharesTable, userShareTable)
	_, err := r.db.Exec(query, userId, id)
	return err
}
