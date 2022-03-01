package http

import (
	"github.com/breeders-zone/auth-service/pkg/jwk"
	"github.com/gofiber/fiber/v2"
)

type JwkResponse struct {
	Keys []interface{} `json:"keys"`
}

func (h *Handler) Jwk(c *fiber.Ctx) error {
	key, err := jwk.GetKey()
	if err != nil {
		c.SendStatus(500)
		return c.SendString("pizdec")
	}

	return c.JSON(&JwkResponse{
		Keys: []interface{}{key},
	})
}
