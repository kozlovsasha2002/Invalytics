package repository

import (
	"database/sql"
	"fmt"
)

const (
	usersTable       = "users"
	depositsTable    = "deposits"
	userDepositTable = "users_deposits"
	bondsTable       = "bonds"
	userBondTable    = "users_bonds"
	sharesTable      = "shares"
	userShareTable   = "users_shares"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	//check connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
