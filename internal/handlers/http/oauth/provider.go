package oauth

import (
	"time"

	"github.com/breeders-zone/auth-service/internal/config"
	"github.com/breeders-zone/auth-service/internal/handlers/http/errors"
	"github.com/breeders-zone/auth-service/internal/services"
	"github.com/breeders-zone/auth-service/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

// OauthProviderRedirect
// @Summary OAuth provider redirect
// @Tags oauth
// @Description Redirect to OAuth provider
// @ModuleID OAuth
// @Accept  json
// @Success 200
// @Param provider path string true "Oauth provider"
// @Router /oauth/{provider} [get]
func (h *Handler) OauthProviderRedirect(c *fiber.Ctx) error {
	return goth_fiber.BeginAuthHandler(c)
}

// OauthProviderCallback
// @Summary OAuth provider callback
// @Tags oauth
// @Description Callback for OAuth provider
// @ModuleID OAuth
// @Accept  json
// @Success 301
// @Failure 400,401,404,500 {object} errors.ErrorResponse
// @Param provider path string true "Oauth provider"
// @Router /oauth/{provider}/callback [get]
func (h *Handler) OauthProviderCallback(c *fiber.Ctx) error {
	oauthUser, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		return c.Status(400).JSON(&errors.ErrorResponse{Code: 400, Message: "Bad request"})
	}

	user, err := h.services.User.FirstOrCreateByEmail(services.FirstOrCreateByEmailInput{
		Email:    oauthUser.Email,
		Name:     oauthUser.FirstName,
		Surename: oauthUser.LastName,
	})

	if err != nil {
		return c.Status(401).JSON(&errors.ErrorResponse{Code: 401, Message: "User not found"})
	}

	token, err := jwt.Create(time.Second*17000000, user.Id)

	if err != nil {
		return c.Status(500).JSON(&errors.ErrorResponse{Code: 500, Message: "Failed to create token"})
	}

	conf, err := config.GetConfig()

	if err != nil {
		return c.Status(500).JSON(&errors.ErrorResponse{Code: 500, Message: "Failed to read config"})
	}

	query := "?"

	if user.Verified {
		query += "verified=1&access_token=" + token
	} else {
		query += "verified=0"
	}

	return c.Redirect(conf.ClientCallbackUrl + query)
}
