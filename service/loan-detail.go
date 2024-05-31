package service

import (
	"context"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"log"
)

//func (s *AppService) UpdateLoanDetail(ctx context.Context, detail model.LoanDetail) error {
//	if err := s.validator.Struct(detail); err != nil {
//		log.Printf("error when validate request %v\n", err)
//		return err
//	}
//
//	ld, err := s.repo.GetLoanDetail(ctx, detail.ID)
//	if err != nil {
//		log.Printf("error when get loan detail %v\n", err)
//		return err
//	}
//	if ld.ID == 0 {
//		return errors.New("loan general not exist")
//	}
//	if ld.Status == 1 {
//		return errors.New("paid loan cannot be updated")
//	}
//	if detail.Amount == ld.Amount {
//		return nil
//	}
//
//	// check loan general
//	lg, err := s.repo.GetLoanGeneral(ctx, detail.LoanGeneralID)
//	if err != nil {
//		log.Printf("error when get loan general %v\n", err)
//		return err
//	}
//	if lg.ID == 0 {
//		return errors.New("loan general not exist")
//	}
//
//	// contains transaction to update remaining loan
//	if err := s.repo.UpdateLoanDetail(ctx, detail); err != nil {
//		log.Printf("error when update to database %v\n", err)
//		return err
//	}
//
//	return nil
//}

func (s *AppService) DeleteLoanDetail(ctx context.Context, id, lgid uint32) error {
	// update tenor and remaining loan
	if err := s.repo.DeleteLoanDetail(ctx, id, lgid); err != nil {
		log.Printf("error when delete from database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) GetLoanDetail(ctx context.Context, id uint32) (model.LoanDetail, error) {
	ld, err := s.repo.GetLoanDetail(ctx, id)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return model.LoanDetail{}, err
	}

	lg, err := s.repo.GetLoanGeneral(ctx, ld.LoanGeneralID)
	if err != nil {
		log.Printf("error when get loan type %v\n", err)
		return model.LoanDetail{}, err
	}
	if lg.ID != 0 {
		ld.General = model.GeneralDetail{
			Title:    lg.Title,
			Amount:   lg.Amount,
			Datetime: lg.Datetime,
			Tenor:    lg.Tenor,
			Status:   lg.Status,
		}
	}

	return ld, nil
}

func (s *AppService) GetLoanDetails(ctx context.Context, generalID uint32) ([]model.LoanDetail, error) {
	lds, err := s.repo.GetLoanDetails(ctx, generalID, nil)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return nil, err
	}

	return lds, nil
}

func (s *AppService) GetMonthlyLoanDetails(ctx context.Context, month int) ([]model.LoanDetail, error) {
	lds, err := s.repo.GetMonthlyLoanDetail(ctx, month)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return nil, err
	}

	return lds, nil
}
