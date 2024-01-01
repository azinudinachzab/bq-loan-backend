package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/azinudinachzab/bq-loan-be-v2/model"
)

func (s *AppService) Registration(ctx context.Context, req model.RegistrationRequest) error {
	if err := s.validator.Struct(req); err != nil {
		log.Printf("error when validate request %v\n", err)
		return err
	}

	okEmail, err := s.repo.IsEmailExists(ctx, req.Email)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		log.Printf("error when check email %v\n", err)
		return err
	}

	if okEmail {
		log.Printf("email is exists %v\n", req.Email)
		return errors.New("email already exist")
	}

	if req.Role == 0 {
		req.Role = 1
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error when hash password %v\n", err)
		return err
	}
	req.Password = string(hashed)

	if err := s.repo.StoreUser(ctx, req); err != nil {
		log.Printf("error when store user %v\n", err)
		return err
	}

	return nil
}

func (s *AppService) Login(ctx context.Context, req model.LoginRequest) (model.User, error) {
	if err := s.validator.Struct(req); err != nil {
		log.Printf("error when validate request %v\n", err)
		return model.User{}, err
	}

	usr, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("error when get user %v", err)
		return model.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(req.Password)); err != nil {
		log.Printf("error when check password %v", err)
		return model.User{}, err
	}

	if usr.IsActive == 0 {
		return model.User{}, errors.New("user not active")
	}

	usr.Password = ""
	return usr, nil
}

func (s *AppService) ChangePassword(ctx context.Context, id uint32, oldPw, newPw string) error {
	if oldPw == "" || newPw == "" {
		err := errors.New("password empty")
		log.Printf("error when validate request %v\n", err)
		return err
	}

	usr, err := s.repo.GetUser(ctx, id)
	if err != nil {
		log.Printf("error when get user %v", err)
		return err
	}

	if usr.IsActive == 0 {
		return errors.New("user not active")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(oldPw)); err != nil {
		log.Printf("error when check old password %v", err)
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPw), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error when hash password %v\n", err)
		return err
	}
	newPw = string(hashed)

	if err := s.repo.UpdateUserPassword(ctx, id, newPw); err != nil {
		log.Printf("error when update password %v\n", err)
		return err
	}

	return nil
}
