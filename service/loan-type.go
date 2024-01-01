package service

import (
	"context"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"log"
)

func (s *AppService) CreateLoanType(ctx context.Context, lt model.LoanType) error {
	if err := s.validator.Struct(lt); err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}

	if err := s.repo.CreateLoanType(ctx, lt); err != nil {
		log.Printf("error when store to database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) UpdateLoanType(ctx context.Context, lt model.LoanType) error {
	if err := s.validator.Struct(lt); err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}

	if err := s.repo.UpdateLoanType(ctx, lt.ID, lt); err != nil {
		log.Printf("error when update to database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) DeleteLoanType(ctx context.Context, id uint32) error {
	if err := s.repo.DeleteLoanType(ctx, id); err != nil {
		log.Printf("error when delete from database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) GetLoanType(ctx context.Context, id uint32) (model.LoanType, error) {
	lt, err := s.repo.GetLoanType(ctx, id)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return model.LoanType{}, err
	}

	return lt, nil
}

func (s *AppService) GetLoanTypes(ctx context.Context) ([]model.LoanType, error) {
	lt, err := s.repo.GetLoanTypes(ctx)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return nil, err
	}

	return lt, nil
}
