package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/breeders-zone/auth-service/internal/domain"
	"github.com/breeders-zone/auth-service/internal/services"
	"github.com/breeders-zone/auth-service/pkg/api"
	mockApi "github.com/breeders-zone/auth-service/pkg/api/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var errInternalServErr = errors.New("test: internal server error")

func mockUserService(t *testing.T) (*services.User, *mockApi.MockAuthServiceClient) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockGreeterClient := mockApi.NewMockAuthServiceClient(mockCtl)

	userService := services.NewUserService(
		mockGreeterClient,
	)

	return userService, mockGreeterClient
}

func TestNewUserService_LoginErr(t *testing.T) {
	userService, mockGreeterClient := mockUserService(t)

	ctx := context.Background()

	mockGreeterClient.EXPECT().Login(ctx, gomock.Any()).Return(&api.UserResponse{}, errInternalServErr)

	res, err := userService.Login(services.UserLoginInput{})

	require.True(t, errors.Is(err, errInternalServErr))
	require.Equal(t, domain.User{}, res)
}

func TestNewUserService_Login(t *testing.T) {
	userService, mockGreeterClient := mockUserService(t)

	ctx := context.Background()

	mockGreeterClient.EXPECT().Login(ctx, gomock.Any()).Return(&api.UserResponse{}, nil)


	res, err := userService.Login(services.UserLoginInput{})

	require.NoError(t, err)
	require.Equal(t, domain.User{}, res)
}


func TestNewUserService_FirstOrCreateByEmailErr(t *testing.T) {
	userService, mockGreeterClient := mockUserService(t)

	ctx := context.Background()

	mockGreeterClient.EXPECT().FirstOrCreateByEmail(ctx, gomock.Any()).Return(&api.UserResponse{}, errInternalServErr)

	res, err := userService.FirstOrCreateByEmail(services.FirstOrCreateByEmailInput{})

	require.True(t, errors.Is(err, errInternalServErr))
	require.Equal(t, domain.User{}, res)
}

func TestNewUserService_FirstOrCreateByEmail(t *testing.T) {
	userService, mockGreeterClient := mockUserService(t)

	ctx := context.Background()

	mockGreeterClient.EXPECT().FirstOrCreateByEmail(ctx, gomock.Any()).Return(&api.UserResponse{}, nil)

	res, err := userService.FirstOrCreateByEmail(services.FirstOrCreateByEmailInput{})

	require.NoError(t, err)
	require.Equal(t, domain.User{}, res)
}