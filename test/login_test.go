package tests

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"

	"github.com/breeders-zone/auth-service/pkg/api"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	unavailableErr = status.Error(codes.Unavailable, "User Service not available")
	internalErr = status.Error(codes.Internal, "User Service not available")
)

func (s *APITestSuite) TestUserLogin() {
	app := fiber.New()

	s.handler.Init(app)
	
	ctx := context.Background()
	s.mocks.apiService.EXPECT().Login(ctx, gomock.Any()).Return(&api.UserResponse{Id: 5}, nil)
	phone, passowrd := "+799999999", "secret"
	signUpData := fmt.Sprintf(`{"phone":"%s", "password":"%s"}`, phone, passowrd)
	
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte(signUpData)))
	req.Header.Add("Content-Type", "application/json")
	

	res, err := app.Test(req, -1)

	
	s.NoError(err)
	s.Equal(res.StatusCode, 200)
}

func (s *APITestSuite) TestUserLoginErrorServiceUnavalible() {
	app := fiber.New()
	s.handler.Init(app)
	
	ctx := context.Background()
	s.mocks.apiService.EXPECT().Login(ctx, gomock.Any()).Return(&api.UserResponse{}, unavailableErr)
	phone, passowrd := "+799999999", "secret"
	signUpData := fmt.Sprintf(`{"phone":"%s", "password":"%s"}`, phone, passowrd)
	
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte(signUpData)))
	req.Header.Add("Content-Type", "application/json")
	

	res, _ := app.Test(req)
	
	s.Equal(res.StatusCode, 503)
}

func (s *APITestSuite) TestUserLoginErrorServiceIternal() {
	app := fiber.New()
	s.handler.Init(app)
	
	ctx := context.Background()
	s.mocks.apiService.EXPECT().Login(ctx, gomock.Any()).Return(&api.UserResponse{}, internalErr)
	phone, passowrd := "+799999999", "secret"
	signUpData := fmt.Sprintf(`{"phone":"%s", "password":"%s"}`, phone, passowrd)
	
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte(signUpData)))
	req.Header.Add("Content-Type", "application/json")
	

	res, _ := app.Test(req)
	
	s.Equal(res.StatusCode, 500)
}