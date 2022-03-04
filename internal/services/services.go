package services

import (
	"github.com/breeders-zone/auth-service/internal/domain"
	"github.com/breeders-zone/auth-service/pkg/api"
)

type UserService interface {
	Login(input UserLoginInput) (domain.User, error)
	FirstOrCreateByEmail(input FirstOrCreateByEmailInput) (domain.User, error)
}

type Services struct {
	User UserService
}

func NewServices(authService api.AuthServiceClient) *Services  {
	return &Services{
		User: NewUserService(authService),
	}
}
