package http

import (
	"github.com/breeders-zone/auth-service/internal/handlers/http/errors"
	"github.com/gofiber/fiber/v2"
)

type JwkResponse struct {
	Keys []interface{} `json:"keys"`
}

func (h *Handler) Jwk(c *fiber.Ctx) error {
	key, err :=  h.tokenManager.GetJwk()
	if err != nil {
		return c.Status(500).JSON(&errors.ErrorResponse{500, "Can't create jwk"})
	}

	return c.JSON(&JwkResponse{
		Keys: []interface{}{key},
	})
}
