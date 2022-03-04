package oauth

import (
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

func (h Handler) Init(router fiber.Router) {
	oauth := router.Group("oauth")

	oauth.Get("/:provider", h.OauthProviderRedirect)
	oauth.Get("/:provider/callback", h.OauthProviderCallback)
}
