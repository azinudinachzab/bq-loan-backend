package service

import (
	"context"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"github.com/azinudinachzab/bq-loan-be-v2/repository"
	"github.com/go-playground/validator/v10"
)

type Dependency struct {
	Validator *validator.Validate
	Repo      repository.Repository
	Conf      model.Configuration
}

type AppService struct {
	validator *validator.Validate
	repo      repository.Repository
	conf      model.Configuration
}

func NewService(dep Dependency) *AppService {
	return &AppService{
		validator: dep.Validator,
		repo:      dep.Repo,
		conf:      dep.Conf,
	}
}

type Service interface {
	Dashboard(ctx context.Context) (model.Dashboard, error)
	DashboardAdmin(ctx context.Context) (model.DashboardAdmin, error)

	Registration(ctx context.Context, req model.RegistrationRequest) error
	Login(ctx context.Context, req model.LoginRequest) (model.User, error)
	ChangePassword(ctx context.Context, id uint32, oldPw, newPw string) error

	ToggleIsActive(ctx context.Context, id uint32) error
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, id uint32) error
	GetUser(ctx context.Context, id uint32) (model.User, error)
	GetUsers(ctx context.Context, filter map[string]string) ([]model.User, error)

	CreateLoanType(ctx context.Context, lt model.LoanType) error
	UpdateLoanType(ctx context.Context, lt model.LoanType) error
	DeleteLoanType(ctx context.Context, id uint32) error
	GetLoanType(ctx context.Context, id uint32) (model.LoanType, error)
	GetLoanTypes(ctx context.Context) ([]model.LoanType, error)

	CreateLoanGeneral(ctx context.Context, general model.LoanGeneral) error
	UpdateLoanGeneral(ctx context.Context, general model.LoanGeneral) error
	DeleteLoanGeneral(ctx context.Context, id uint32) error
	GetLoanGeneral(ctx context.Context, id uint32) (model.LoanGeneral, error)
	GetLoanGenerals(ctx context.Context, lastID, uID uint32, title string) ([]model.LoanGeneral, error)

	AcceptLoanRequest(ctx context.Context, id uint32) error
	AcceptPaymentRequest(ctx context.Context, id uint32) error

	//UpdateLoanDetail(ctx context.Context, detail model.LoanDetail) error
	DeleteLoanDetail(ctx context.Context, id, lgid uint32) error
	GetLoanDetail(ctx context.Context, id uint32) (model.LoanDetail, error)
	GetLoanDetails(ctx context.Context, generalID uint32) ([]model.LoanDetail, error)

	SocialFundRequest(ctx context.Context, req model.SocialFund) error
	AcceptSocialFundRequest(ctx context.Context, id uint32) error
}
