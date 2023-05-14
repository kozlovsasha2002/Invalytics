package service

import (
	"Invalytics/app/internal/dto"
	"Invalytics/app/internal/model"
	"Invalytics/app/internal/repository"
)

type Service struct {
	Authorization
	Deposit
	Bond
	Share
	Profit
	Company
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Deposit:       NewDepositService(repo),
		Bond:          NewBondService(repo),
		Share:         NewShareService(repo),
		Profit:        NewProfitService(repo),
		Company:       NewCompanyService(repo),
	}
}

type Authorization interface {
	CreateUser(user model.User) (int32, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int32, error)
}

type Deposit interface {
	CreateDeposit(userId int32, deposit model.Deposit) (int32, error)
	GetAllDeposits(userId int32) ([]model.Deposit, error)
	GetDepositById(userId, id int32) (model.Deposit, error)
	UpdateDeposit(userId, id int32, deposit model.UpdateDeposit) error
	DeleteDeposit(userId, id int32) error
}

type Bond interface {
	CreateBond(userId int32, bond model.Bond) (int32, error)
	GetAllBonds(userId int32) ([]model.Bond, error)
	GetBondById(userId, id int32) (model.Bond, error)
	UpdateBond(userId, id int32, input model.UpdateBond) error
	DeleteBond(userId, id int32) error
}

type Share interface {
	CreateShare(userId int32, share model.Share) (int32, error)
	GetAllShares(userId int32) ([]model.Share, error)
	GetShareById(userId, id int32) (model.Share, error)
	UpdateShare(userId, id int32, input model.UpdateShare) error
	DeleteShare(userId, id int32) error
}

type Profit interface {
	ShareProfitabilityById(userId, shareId int32, termInMonths int) (model.ProfitInfo, error)
	AllShareProfitability(userId int32, termInMonths int, sort bool) ([]model.ProfitInfo, error)
	BondProfitabilityById(userId, bondId int32) (dto.BondInfo, error)
	AllBondProfitability(userId int32, srt bool) ([]dto.BondInfo, error)
	DepositProfitabilityById(userId, depositId int32) (dto.DepositDto, error)
	AllDepositProfitability(userId int32, srt bool) ([]dto.DepositDto, error)
}

type Company interface {
	CreateCompany(userId int32, comp model.Company) (int32, error)
	GetAllCompanies(userId int32) ([]model.Company, error)
	GetCompanyById(userId, compId int32) (model.Company, error)
	UpdateCompany(userId, compId int32, input dto.UpdateCompanyDto) error
	DeleteCompany(userId, compId int32) error
	GetAllMultipliers(userId int32, param string) ([]dto.MultiplierDto, error)
	GetMultiplierById(userId, compId int32) (dto.MultiplierDto, error)
}
