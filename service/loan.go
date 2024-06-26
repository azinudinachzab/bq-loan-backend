package service

import (
	"context"
	"errors"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"log"
	"math"
	"time"
)

func (s *AppService) AcceptLoanRequest(ctx context.Context, id uint32) error {
	lg, err := s.repo.GetLoanGeneral(ctx, id)
	if err != nil {
		log.Printf("error when get loan general %v\n", err)
		return err
	}

	if lg.Status == 1 {
		return errors.New("already paid")
	}

	lt, err := s.repo.GetLoanType(ctx, lg.LoanTypeID)
	if err != nil {
		log.Printf("error when get loan type %v\n", err)
		return err
	}

	usr, err := s.repo.GetUser(ctx, lg.UserID)
	if err != nil {
		log.Printf("error when get user %v\n", err)
		return err
	}

	if lt.IsAddBalance != 1 {
		switch lt.IsAddBalance {
		case 2:
			if err := s.repo.AddBalance(ctx, lg.ID, lg.UserID, "balance", usr.Balance, lg.Amount); err != nil {
				return err
			}
			return nil
		case 3:
			if err := s.repo.AddBalance(ctx, lg.ID, lg.UserID, "vbalance", usr.VBalance, lg.Amount); err != nil {
				return err
			}
			return nil
		case 4:
			if err := s.repo.AddBalance(ctx, lg.ID, lg.UserID, "social", 0, lg.Amount); err != nil {
				return err
			}
			return nil
		}
	}

	loanDetails := make([]model.LoanDetail, 0)
	tmp := model.LoanDetail{
		LoanGeneralID: lg.ID,
		Amount:        lg.Amount / float64(lg.Tenor),
		Status:        0,
	}

	now := time.Now()
	for i := 0; i < lg.Tenor; i++ {
		date := time.Date(now.Year(), now.Month(), 25, now.Hour(), now.Minute(), now.Second(),
			now.Nanosecond(), now.Location()).AddDate(0, i+1, 0)
		tmp.Datetime = date.Format(time.DateTime)

		loanDetails = append(loanDetails, tmp)
	}

	if len(loanDetails) > 0 {
		err = s.repo.BulkSaveLoanDetail(ctx, loanDetails)
		if err != nil {
			log.Printf("error when bulk save %v\n", err)
			return err
		}
	}

	return nil
}

func (s *AppService) AcceptPaymentRequest(ctx context.Context, id uint32) error {
	ld, err := s.repo.GetLoanDetail(ctx, id)
	if err != nil {
		log.Printf("error when get loan detail %v\n", err)
		return err
	}

	if ld.Status == 1 {
		return errors.New("already paid")
	}

	lg, err := s.repo.GetLoanGeneral(ctx, ld.LoanGeneralID)
	if err != nil {
		log.Printf("error when get loan general %v\n", err)
		return err
	}

	lt, err := s.repo.GetLoanType(ctx, lg.LoanTypeID)
	if err != nil {
		log.Printf("error when get loan type %v\n", err)
		return err
	}

	incomeAmount := 0.00
	if lt.Margin != 0 {
		percentage := 100.00 + lt.Margin
		originalAmount := (lg.Amount / percentage) * 100.00
		incomeAmount = math.Round(originalAmount / lt.Margin / float64(lg.Tenor))
	}

	return s.repo.UpdateLoanDetailStatus(ctx, ld.ID, ld.LoanGeneralID, incomeAmount)
}

func (s *AppService) Dashboard(ctx context.Context) (model.Dashboard, error) {
	return s.repo.GetDashboardData(ctx)
}

func (s *AppService) DashboardAdmin(ctx context.Context) (model.DashboardAdmin, error) {
	return s.repo.GetDashboardAdminData(ctx)
}

func (s *AppService) SocialFundRequest(ctx context.Context, req model.SocialFund) error {
	if err := s.validator.Struct(req); err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}

	usr, err := s.repo.GetUser(ctx, req.UserID)
	if err != nil {
		log.Printf("error when get user %v\n", err)
		return err
	}
	if usr.ID == 0 {
		return errors.New("user not registered")
	}

	if err := s.repo.CreateSocialFundRequest(ctx, req); err != nil {
		log.Printf("error when store social fund %v\n", err)
		return err
	}
	return nil
}

func (s *AppService) AcceptSocialFundRequest(ctx context.Context, id uint32) error {
	return s.repo.UpdateSocialFundStatus(ctx, id)
}
