package http

import (
	"fmt"
	"time"

	"github.com/breeders-zone/auth-service/internal/domain"
	"github.com/breeders-zone/auth-service/internal/services"
	"github.com/go-playground/validator/v10"

	"github.com/breeders-zone/auth-service/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string       `json:"access_token"`
	User  *domain.User `json:"data"`
}

// Login
// @Summary User Login
// @Tags users-auth
// @Description Login user
// @ModuleID SignUp
// @Accept  json
// @Produce  json
// @Param input body LoginRequest true "sign in info"
// @Success 200 {object} LoginResponse
// @Failure 400,404 {object} LoginRequest
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} LoginRequest
// @Router /login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	input := new(LoginRequest)
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	var errors []*ErrorResponse
	if err := validator.New().Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	if errors != nil {
		return c.Status(422).JSON(errors)
	}

	user, err := h.services.User.Login(services.UserLoginInput{
		Phone:    input.Phone,
		Password: input.Password,
	})

	if err != nil {
		fmt.Print(err)
		return c.JSON("User not found")
	}

	token, _ := jwt.Create(time.Second*17000000, user.Id)

	return c.JSON(&LoginResponse{
		token,
		user,
	})
}
