package services

import "github.com/breeders-zone/auth-service/pkg/api"

type Services struct {
	User UserService
}

func NewServices(authService api.AuthServiceClient) *Services  {
	return &Services{
		User: NewUserService(authService),
	}
}
