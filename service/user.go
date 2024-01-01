package service

import (
	"context"
	"errors"
	"log"

	"github.com/azinudinachzab/bq-loan-be-v2/model"
)

func (s *AppService) ToggleIsActive(ctx context.Context, id uint32) error {
	usr, err := s.repo.GetUser(ctx, id)
	if err != nil {
		log.Printf("error when get user %v\n", err)
		return err
	}

	mapToggle := map[int]int{
		0: 1,
		1: 0,
	}
	toggleIsActive, ok := mapToggle[usr.IsActive]
	if !ok {
		return errors.New("active status not recognized")
	}

	if err := s.repo.UpdateIsActive(ctx, id, toggleIsActive); err != nil {
		log.Printf("error when update activation %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) UpdateUser(ctx context.Context, user model.User) error {
	sTemp := struct {
		Name     string  `validate:"required"`
		Email    string  `validate:"required"`
		Role     int     `validate:"required"`
		IsActive int     `validate:"required"`
		IsLeader int     `validate:"required"`
		Balance  float64 `validate:"required"`
		VBalance float64 `validate:"required"`
	}{
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		IsActive: user.IsActive,
		IsLeader: user.IsLeader,
		Balance:  user.Balance,
		VBalance: user.VBalance,
	}
	if err := s.validator.Struct(sTemp); err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}
	if err := s.repo.UpdateUser(ctx, user.ID, user); err != nil {
		log.Printf("error when update to database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) DeleteUser(ctx context.Context, id uint32) error {
	if err := s.repo.DeleteUser(ctx, id); err != nil {
		log.Printf("error when delete from database %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) GetUser(ctx context.Context, id uint32) (model.User, error) {
	usr, err := s.repo.GetUser(ctx, id)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return model.User{}, err
	}

	usr.Password = ""

	return usr, nil
}

func (s *AppService) GetUsers(ctx context.Context, filter map[string]string) ([]model.User, error) {
	usr, err := s.repo.GetUsersByFilter(ctx, filter)
	if err != nil {
		log.Printf("error when get data from database %v\n", err)
		return nil, err
	}

	return usr, nil
}
