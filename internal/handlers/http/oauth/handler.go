package oauth

import (
	"github.com/breeders-zone/auth-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services,
	}
}

func (h Handler) Init(router fiber.Router) {
	oauth := router.Group("oauth")

	oauth.Get("/:provider", h.OauthProviderRedirect)
	oauth.Get("/:provider/callback", h.OauthProviderCallback)
}
