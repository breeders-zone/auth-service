package http

import (
	"time"

	"github.com/breeders-zone/auth-service/internal/domain"
	"github.com/breeders-zone/auth-service/internal/handlers/http/errors"
	"github.com/breeders-zone/auth-service/internal/services"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string       `json:"access_token"`
	User  domain.User `json:"data"`
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
// @Failure 400,401,500,503 {object} errors.ErrorResponse
// @Failure 422 {object} errors.ValidationErrorResponse
// @Failure default {object} LoginRequest
// @Router /login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	input := new(LoginRequest)
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	var valErrors []errors.ValidationError
	if err := validator.New().Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errors.ValidationError
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			valErrors = append(valErrors, element)
		}
	}

	if valErrors != nil {
		return c.Status(422).JSON(errors.ValidationErrorResponse{
			ErrorResponse: errors.ErrorResponse{
				Code: 422,
				 Message: "Incalid request",
			}, 
			Errors: valErrors,
		})
	}

	user, err := h.services.User.Login(services.UserLoginInput{
		Phone:    input.Phone,
		Password: input.Password,
	})

	if err != nil {
		if e, ok := status.FromError(err); ok {
            switch e.Code() {
				case codes.Unavailable:
					return c.Status(503).JSON(&errors.ErrorResponse{Code: 503, Message: "User Service not available"})
				case codes.Internal:
					return c.Status(500).JSON(&errors.ErrorResponse{Code: 500, Message: "User Service internal server error"})
            }
        }


		return c.Status(401).JSON(&errors.ErrorResponse{Code: 401, Message: "User not found"})
	}

	token, err := h.tokenManager.Create(time.Second*17000000, user.Id, []string{"user"})

	if err != nil {
		return c.Status(500).JSON(&errors.ErrorResponse{Code: 500, Message: "Failed to create token"})
	}

	return c.JSON(&LoginResponse{
		token,
		user,
	})
}
