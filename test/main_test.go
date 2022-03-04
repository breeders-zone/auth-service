package tests

import (
	"log"
	"os"
	"testing"

	"github.com/breeders-zone/auth-service/internal/handlers/http"
	"github.com/breeders-zone/auth-service/internal/services"
	mockApi "github.com/breeders-zone/auth-service/pkg/api/mocks"
	"github.com/breeders-zone/auth-service/pkg/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite

	handler  *http.Handler
	services *services.Services

	mocks *mocks
}

type mocks struct {
	apiService *mockApi.MockAuthServiceClient
}

func init() {

	env := struct {
		OauthBase         string
		ClientCallbackUrl string
		JwtPublicKey      string
		JwtPrivateKey     string
		JwtKeyId          string
		AuthGrpcServer    string
		VkCleintId        string
		VkClientSercet    string
	}{
		OauthBase:         "http://localhost:3000/oauth/",
		ClientCallbackUrl: "http://localhost/auth/callback",
		JwtPublicKey:      "./data/certs/is_rsa.pub",
		JwtPrivateKey:     "./data/certs/is_rsa",
		JwtKeyId:          "qwerty",
		AuthGrpcServer:    ":9002",
		VkCleintId:        "12312",
		VkClientSercet:    "secret",
	}

	os.Setenv("OAUTH_BASE", env.OauthBase)
	os.Setenv("CLIENT_CALLBACK_URL", env.ClientCallbackUrl)
	os.Setenv("AUTH_GRPC_SERVER", env.AuthGrpcServer)
	os.Setenv("JWT_PUBLIC_KEY", env.JwtPublicKey)
	os.Setenv("JWT_PRIVATE_KEY", env.JwtPrivateKey)
	os.Setenv("JWT_KEY_ID", env.JwtKeyId)
	os.Setenv("VKONTAKTE_CLIENT_ID", env.VkCleintId)
	os.Setenv("VKONTAKTE_CLIENT_SECRET", env.VkClientSercet)
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupSuite() {

	s.initMocks()
	s.initDeps()
}

func (s *APITestSuite) initMocks() {
	mockCtl := gomock.NewController(s.T())
	defer mockCtl.Finish()

	s.mocks = &mocks{
		apiService: mockApi.NewMockAuthServiceClient(mockCtl),
	}
}

func (s *APITestSuite) initDeps() {
	tokenManager, err := auth.NewTokenManager("token", "", "")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	services := services.NewServices(s.mocks.apiService)
	s.handler = http.NewHandler(services, tokenManager)
	s.services = services
}

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}
