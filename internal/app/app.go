package app

import (
	"log"

	"github.com/breeders-zone/auth-service/internal/config"
	"github.com/breeders-zone/auth-service/internal/handlers/http"
	"github.com/breeders-zone/auth-service/internal/services"
	"github.com/breeders-zone/auth-service/pkg/api"
	"github.com/breeders-zone/auth-service/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/vk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Breeders Zone Auth API
// @version 1.0
// @description REST API for breeders zone auth service

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

// Run initializes whole application.
func Run() {
	app := fiber.New()

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Not load config")
	}


	conn, err := grpc.Dial(conf.AuthGrpcServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	authService := api.NewAuthServiceClient(conn)

	goth.UseProviders(
		vk.New(conf.VkCleintId, conf.VkClientSercet, conf.OauthBase + "vk/callback"),
	)

	tokenManager, err := auth.NewTokenManager(conf.JwtKeyId, conf.JwtPublicKey, conf.JwtPrivateKey)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	services := services.NewServices(authService)
	handler := http.NewHandler(services, tokenManager)
	handler.Init(app)

	app.Listen(":3000")
}
