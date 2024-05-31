package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"log"
)

type PgRepository struct {
	dbCore *sql.DB
}

func NewPgRepository(dbCore *sql.DB) *PgRepository {
	return &PgRepository{
		dbCore: dbCore,
	}
}

func rowClose(r *sql.Rows) {
	if err := r.Close(); err != nil {
		log.Printf("error when closing row %v\n", err)
	}
}

type Repository interface {
	GetDashboardData(ctx context.Context) (model.Dashboard, error)
	GetDashboardAdminData(ctx context.Context) (model.DashboardAdmin, error)

	CreateLoanType(ctx context.Context, loanType model.LoanType) error
	GetLoanTypes(ctx context.Context) ([]model.LoanType, error)
	GetLoanType(ctx context.Context, id uint32) (model.LoanType, error)
	UpdateLoanType(ctx context.Context, id uint32, lt model.LoanType) error
	DeleteLoanType(ctx context.Context, id uint32) error

	IsEmailExists(ctx context.Context, email string) (bool, error)
	StoreUser(ctx context.Context, regData model.RegistrationRequest) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUser(ctx context.Context, id uint32) (model.User, error)
	UpdateIsActive(ctx context.Context, id uint32, isActive int) error
	GetUsersByFilter(ctx context.Context, filter map[string]string) ([]model.User, error)
	UpdateUser(ctx context.Context, id uint32, usr model.User) error
	DeleteUser(ctx context.Context, id uint32) error
	UpdateUserPassword(ctx context.Context, id uint32, newPw string) error

	CreateLoanGeneral(ctx context.Context, general model.LoanGeneral) error
	UpdateLoanGeneral(ctx context.Context, general model.LoanGeneral) error
	DeleteLoanGeneral(ctx context.Context, id uint32) error
	GetLoanGeneral(ctx context.Context, id uint32) (model.LoanGeneral, error)
	GetLoanGenerals(ctx context.Context, lastID, uID uint32, title string) ([]model.LoanGeneral, error)

	//UpdateLoanDetail(ctx context.Context, detail model.LoanDetail) error
	DeleteLoanDetail(ctx context.Context, id, lgid uint32) error
	GetLoanDetail(ctx context.Context, id uint32) (model.LoanDetail, error)
	GetLoanDetails(ctx context.Context, generalID uint32, tx *sql.Tx) ([]model.LoanDetail, error)
	UpdateLoanDetailStatus(ctx context.Context, id, lgid uint32, amount float64) error
	BulkSaveLoanDetail(ctx context.Context, dt []model.LoanDetail) error
	AddBalance(ctx context.Context, id, uid uint32, balance string, c, amount float64) error
	GetMonthlyLoanDetail(ctx context.Context, month int) ([]model.LoanDetail, error)

	CreateSocialFundRequest(ctx context.Context, sf model.SocialFund) error
	UpdateSocialFundStatus(ctx context.Context, id uint32) error
}

func (p *PgRepository) GetDashboardData(ctx context.Context) (model.Dashboard, error) {
	q := `SELECT (SELECT amount FROM loan_generals WHERE status=1 AND (loan_type_id=3 OR loan_type_id=4) ORDER BY datetime DESC LIMIT 1) AS latest_loan,
       (SELECT amount FROM loan_generals WHERE status=1 AND (loan_type_id=3 OR loan_type_id=4) ORDER BY amount DESC LIMIT 1) AS biggest_loan,
       (SELECT count(id) FROM loan_generals WHERE status=1 AND (loan_type_id=3 OR loan_type_id=4)) AS accepted_loan`
	var (
		ll, bl sql.NullFloat64
		al     sql.NullInt32
	)
	if err := p.dbCore.QueryRowContext(ctx, q).Scan(&ll, &bl, &al); err != nil {
		return model.Dashboard{}, err
	}
	return model.Dashboard{
		LatestLoan:   ll.Float64,
		BiggestLoan:  bl.Float64,
		AcceptedLoan: int(al.Int32),
	}, nil
}

func (p *PgRepository) GetDashboardAdminData(ctx context.Context) (model.DashboardAdmin, error) {
	q := `SELECT (SELECT SUM(amount) FROM income) AS total_income`
	var (
		ti sql.NullFloat64
	)
	if err := p.dbCore.QueryRowContext(ctx, q).Scan(&ti); err != nil {
		return model.DashboardAdmin{}, err
	}
	return model.DashboardAdmin{
		TotalIncome: fmt.Sprintf("%.2f", ti.Float64),
	}, nil
}
