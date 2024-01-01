package service

import (
	"context"
	"errors"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"log"
	"time"
)

func (s *AppService) CreateLoanGeneral(ctx context.Context, general model.LoanGeneral) error {
	if err := s.validator.Struct(general); err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}

	t, err := time.Parse(time.RFC3339, general.Datetime)
	if err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}

	// check users
	usr, err := s.repo.GetUser(ctx, general.UserID)
	if err != nil {
		log.Printf("error when get user %v\n", err)
		return err
	}
	if usr.ID == 0 {
		return errors.New("user not registered")
	}

	// check loan type
	lt, err := s.repo.GetLoanType(ctx, general.LoanTypeID)
	if err != nil {
		log.Printf("error when get loan type %v\n", err)
		return err
	}
	if lt.ID == 0 {
		return errors.New("loan type not exist")
	}

	if lt.Margin > 0 {
		general.Amount = general.Amount + ((general.Amount * lt.Margin) / 100)
	}
	general.Status = 0
	general.Datetime = t.Format(time.DateTime)
	if err := s.repo.CreateLoanGeneral(ctx, general); err != nil {
		log.Printf("error when store to database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) UpdateLoanGeneral(ctx context.Context, general model.LoanGeneral) error {
	if err := s.validator.Struct(general); err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}
	t, err := time.Parse(time.RFC3339, general.Datetime)
	if err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}

	// check users
	usr, err := s.repo.GetUser(ctx, general.UserID)
	if err != nil {
		log.Printf("error when get user %v\n", err)
		return err
	}
	if usr.ID == 0 {
		return errors.New("user not registered")
	}

	// check loan type
	lt, err := s.repo.GetLoanType(ctx, general.LoanTypeID)
	if err != nil {
		log.Printf("error when get loan type %v\n", err)
		return err
	}
	if lt.ID == 0 {
		return errors.New("loan type not exist")
	}

	lg, err := s.repo.GetLoanGeneral(ctx, general.ID)
	if err != nil {
		log.Printf("error when get loan general %v\n", err)
		return err
	}

	if lg.Status != 0 {
		return errors.New("cant edit accepted loan")
	}

	if lt.Margin > 0 {
		general.Amount = general.Amount + ((general.Amount * lt.Margin) / 100)
	}

	// contains transaction to update loan detail
	general.Datetime = t.Format(time.DateTime)
	if err := s.repo.UpdateLoanGeneral(ctx, general); err != nil {
		log.Printf("error when update to database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) DeleteLoanGeneral(ctx context.Context, id uint32) error {
	if err := s.repo.DeleteLoanGeneral(ctx, id); err != nil {
		log.Printf("error when delete from database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) GetLoanGeneral(ctx context.Context, id uint32) (model.LoanGeneral, error) {
	lg, err := s.repo.GetLoanGeneral(ctx, id)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return model.LoanGeneral{}, err
	}

	usr, err := s.repo.GetUser(ctx, lg.UserID)
	if err != nil {
		log.Printf("error when get user %v\n", err)
		return model.LoanGeneral{}, err
	}
	if usr.ID != 0 {
		lg.Users = model.UserGeneral{
			ID:   usr.ID,
			Name: usr.Name,
		}
	}

	lt, err := s.repo.GetLoanType(ctx, lg.LoanTypeID)
	if err != nil {
		log.Printf("error when get loan type %v\n", err)
		return model.LoanGeneral{}, err
	}
	if lt.ID != 0 {
		lg.LoanTypes = model.LoanTypeGeneral{
			ID:   lt.ID,
			Name: lt.Name,
		}
	}

	return lg, nil
}

func (s *AppService) GetLoanGenerals(ctx context.Context, lastID, uID uint32, title string) ([]model.LoanGeneral, error) {
	lgs, err := s.repo.GetLoanGenerals(ctx, lastID, uID, title)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return nil, err
	}

	return lgs, nil
}
