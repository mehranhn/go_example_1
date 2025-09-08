package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mehranhn/go_example_1/models/request"
	"github.com/mehranhn/go_example_1/services"
	"github.com/mehranhn/go_example_1/utils"
)

type UserController struct {
	userService services.UserService
	validate    *validator.Validate
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService: userService,
		validate:    validator.New(),
	}
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user details by their UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID" Format(uuid)
// @Success 200 {object} response.UserDto "User details"
// @Failure 400 "Invalid UUID format"
// @Failure 404 "User not found"
// @Failure 500 "Internal server error"
// @Security BearerAuth
// @Router /user/{id} [get]
func (controller *UserController) GetUserByID(c *fiber.Ctx) error {
	id, err := utils.ReadUUIDParamFiber(c, "id")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "Invcalid UUID Param",
		})
	}

	user, err := controller.userService.GetUserById(id)
	if err != nil {
		return c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	if user == nil {
		return c.SendStatus(fiber.ErrNotFound.Code)
	}

	return c.JSON(user)
}

// GetList godoc
// @Summary Get users list
// @Description Get paginated list of users with optional filtering
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1) minimum(1)
// @Param limit query int false "Items per page" default(10) minimum(1) maximum(100)
// @Param search query string false "Filter by phone number"
// @Success 200 {object} []response.UserDto "List of users"
// @Failure 400 "Invalid query parameters"
// @Failure 500 "Internal server error"
// @Security BearerAuth
// @Router /user/ [get]
func (controller *UserController) GetList(c *fiber.Ctx) error {
	var filter request.PaginationFilter

	err := c.QueryParser(&filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	err = controller.validate.Struct(&filter)
	if err != nil {
		return c.SendStatus(fiber.ErrUnprocessableEntity.Code)
	}

	users, err := controller.userService.GetUserList(filter)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(users)
}
