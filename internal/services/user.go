package services

import (
	"context"

	"github.com/breeders-zone/auth-service/internal/domain"
	"github.com/breeders-zone/auth-service/pkg/api"
)

type UserLoginInput struct {
	Phone    string
	Password string
}

type UserService interface {
	Login(input UserLoginInput) (*domain.User, error)
}

type User struct {
	authService api.AuthServiceClient
}

func NewUserService(authService api.AuthServiceClient) *User {
	return &User{
		authService,
	}
}

func (s *User) Login(input UserLoginInput) (*domain.User, error) {

	res, err := s.authService.Login(context.Background(), &api.LoginRequest{Phone: input.Phone, Password: input.Password})
	if err != nil {
		return nil, err
	}

	// some logic
	return &domain.User{
		Id: res.Id,
		Name: res.Name,
		Surname: res.Surname,
		CompanyName: res.CompanyName,
		Phone:    res.Phone,
	}, nil
}
