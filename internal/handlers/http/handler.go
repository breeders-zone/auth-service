package http

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/breeders-zone/auth-service/docs"
	"github.com/breeders-zone/auth-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	app      *fiber.App
	services *services.Services
}

func NewHandler(app *fiber.App, services *services.Services) *Handler {
	return &Handler{
		app,
		services,
	}
}

func (h Handler) Init() {
	h.app.Get("/swagger/*", swagger.HandlerDefault) // default

	h.app.Post("/login", h.Login)
	h.app.Get("/.well-known/jwks.json", h.Jwk)
}
