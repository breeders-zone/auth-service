package http

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/breeders-zone/auth-service/docs"
	"github.com/breeders-zone/auth-service/internal/handlers/http/oauth"
	"github.com/breeders-zone/auth-service/internal/services"
	"github.com/breeders-zone/auth-service/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *services.Services
	tokenManager *auth.TokenManager
}

func NewHandler(services *services.Services, tokenManager *auth.TokenManager) *Handler {
	return &Handler{
		services,
		tokenManager,
	}
}

func (h Handler) Init(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("pong")
	})
	app.Post("/login", h.Login)
	app.Get("/.well-known/jwks.json", h.Jwk)

	oauth.NewHandler(h.services, h.tokenManager).Init(app)
}
