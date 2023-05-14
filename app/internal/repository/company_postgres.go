package repository

import (
	"Invalytics/app/internal/dto"
	"Invalytics/app/internal/model"
	"Invalytics/app/pkg/postgresql"
	"database/sql"
	"fmt"
	"strings"
)

type CompanyPostgres struct {
	db *sql.DB
}

func NewCompanyPostgres(db *sql.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (r *CompanyPostgres) CreateCompany(userId int32, comp model.Company) (int32, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int32
	companyQuery := fmt.Sprintf("INSERT INTO %s (name, dept_payments, depreciation, taxes, market_capitalization, annual_profit, debentures) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", postgresql.CompaniesTable)
	row := tx.QueryRow(companyQuery, comp.Name, comp.DeptPayments, comp.Depreciation, comp.Taxes, comp.MarketCapitalization, comp.AnnualProfit, comp.Debentures)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	userCompanyQuery := fmt.Sprintf("INSERT INTO %s (user_id, company_id) VALUES ($1, $2)", postgresql.UserCompanyTable)
	_, err = tx.Exec(userCompanyQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CompanyPostgres) GetAllCompanies(userId int32) ([]model.Company, error) {
	list := make([]model.Company, 0)
	query := fmt.Sprintf("SELECT comp.id, comp.name, comp.dept_payments, comp.depreciation, comp.taxes, comp.market_capitalization, comp.annual_profit, comp.debentures FROM %s comp INNER JOIN %s usercomp ON comp.id = usercomp.company_id WHERE usercomp.user_id = $1",
		postgresql.CompaniesTable, postgresql.UserCompanyTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c model.Company
		err := rows.Scan(&c.Id, &c.Name, &c.DeptPayments, &c.Depreciation, &c.Taxes, &c.MarketCapitalization, &c.AnnualProfit, &c.Debentures)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

func (r *CompanyPostgres) GetCompanyById(userId, compId int32) (model.Company, error) {
	var c model.Company
	query := fmt.Sprintf("SELECT comp.id, comp.name, comp.dept_payments, comp.depreciation, comp.taxes, comp.market_capitalization, comp.annual_profit, comp.debentures FROM %s comp INNER JOIN %s usercomp ON comp.id = usercomp.company_id WHERE usercomp.user_id = $1 AND usercomp.company_id = $2",
		postgresql.CompaniesTable, postgresql.UserCompanyTable)

	row := r.db.QueryRow(query, userId, compId)
	if err := row.Scan(&c.Id, &c.Name, &c.DeptPayments, &c.Depreciation, &c.Taxes, &c.MarketCapitalization, &c.AnnualProfit, &c.Debentures); err != nil {
		return c, err
	}

	return c, nil
}

func (r *CompanyPostgres) UpdateCompany(userId, compId int32, input dto.UpdateCompanyDto) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.DeptPayments != nil {
		setValues = append(setValues, fmt.Sprintf("dept_payments=$%d", argId))
		args = append(args, *input.DeptPayments)
		argId++
	}

	if input.Depreciation != nil {
		setValues = append(setValues, fmt.Sprintf("depreciation=$%d", argId))
		args = append(args, *input.Depreciation)
		argId++
	}

	if input.Taxes != nil {
		setValues = append(setValues, fmt.Sprintf("taxes=$%d", argId))
		args = append(args, *input.Taxes)
		argId++
	}

	if input.MarketCapitalization != nil {
		setValues = append(setValues, fmt.Sprintf("market_capitalization=$%d", argId))
		args = append(args, *input.MarketCapitalization)
		argId++
	}

	if input.AnnualProfit != nil {
		setValues = append(setValues, fmt.Sprintf("annual_profit=$%d", argId))
		args = append(args, *input.AnnualProfit)
		argId++
	}

	if input.Debentures != nil {
		setValues = append(setValues, fmt.Sprintf("debentures=$%d", argId))
		args = append(args, *input.Debentures)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s comp SET %s FROM %s usercomp WHERE comp.id = usercomp.company_id AND usercomp.user_id=$%d AND usercomp.company_id=$%d",
		postgresql.CompaniesTable, setQuery, postgresql.UserCompanyTable, argId, argId+1)
	args = append(args, userId, compId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *CompanyPostgres) DeleteCompany(userId, compId int32) error {
	query := fmt.Sprintf("DELETE FROM %s comp USING %s usercomp WHERE comp.id = usercomp.company_id AND usercomp.user_id = $1 AND usercomp.company_id = $2",
		postgresql.CompaniesTable, postgresql.UserCompanyTable)
	_, err := r.db.Exec(query, userId, compId)
	return err
}
