// Package controllers
package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mehranhn/go_example_1/constants"
	"github.com/mehranhn/go_example_1/models/request"
	"github.com/mehranhn/go_example_1/services"
)

type AuthController struct {
	authService services.AuthService
	validate    *validator.Validate
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		authService: authService,
		validate:    validator.New(),
	}
}

// RegistryOrLogin godoc
// @Summary Register or login user
// @Description Register a new user or login existing user by phone number. Sends OTP for verification.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.RegisterOrLoginDto true "Registration/Login data"
// @Success 200 "OTP sent successfully and user already existed"
// @Success 201 "OTP sent successfully and reated new user"
// @Failure 400 "Invalid request data"
// @Failure 422 "Validation failed"
// @Failure 500 "Internal server error"
// @Router /auth/register-or-login [post]
func (controller *AuthController) RegistryOrLogin(c *fiber.Ctx) error {
	var dto request.RegisterOrLoginDto
	err := c.BodyParser(&dto)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = controller.validate.Struct(&dto)
	if err != nil {
		return c.SendStatus(fiber.ErrUnprocessableEntity.Code)
	}

	res, err := controller.authService.RegistryOrLogin(dto)
	if err != nil {
		return c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	switch *res {
	case constants.Login:
		return c.SendStatus(fiber.StatusOK);
	case constants.Register:
		return c.SendStatus(fiber.StatusCreated)
	}

	return c.SendStatus(fiber.StatusOK)
}

// ConfirmOtp godoc
// @Summary Confirm OTP code
// @Description Verify OTP code sent to user's phone and complete authentication
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.ConfirmOtpDto true "OTP confirmation data"
// @Success 200 "Authentication successful"
// @Failure 400 "Invalid request data"
// @Failure 401 "Invalid OTP code"
// @Failure 422 "Validation failed"
// @Failure 500 "Internal server error"
// @Router /auth/confirm-otp [post]
func (controller *AuthController) ConfirmOtp(c *fiber.Ctx) error {
	var dto request.ConfirmOtpDto
	err := c.BodyParser(&dto)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = controller.validate.Struct(&dto)
	if err != nil {
		return c.SendStatus(fiber.ErrUnprocessableEntity.Code)
	}

	code, token, _ := controller.authService.ConfirmOtp(dto)
	switch code {
	case 0:
		return c.JSON(token)
	case 1:
		return c.SendStatus(401)
	case 2:
		return c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	return c.SendStatus(fiber.ErrInternalServerError.Code)
}
