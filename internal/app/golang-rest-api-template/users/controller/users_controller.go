package controller

import (
	"context"
	"net/http"

	"github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/dto"
	"github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/models"
	"github.com/gofiber/fiber/v2"
)

// UserService represents the user service interface.
type UserService interface {
	CreateUser(ctx context.Context, user dto.UserCreationDTO) (*models.User, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
}

// UserController represents the user controller.
type UserController struct {
	userService UserService
}

// New creates a new UserController instance.
func New(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// RegisterRoutes registers user-related HTTP routes.
func (uc *UserController) RegisterRoutes(router fiber.Router) {
	router.Get("/users", uc.CreateUser)
	router.Get("/users/{userID}", uc.GetUserByID)
	// Add other routes as needed
}

// CreateUser handles the creation of a new user.
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	createUserDto := new(dto.UserCreationDTO)

	if err := c.BodyParser(createUserDto); err != nil {
		return err
	}

	createdUser, err := uc.userService.CreateUser(c.Context(), *createUserDto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(createdUser)

}

// GetUserByID retrieves a user by their ID.
func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("userID")

	user, err := uc.userService.GetUserByID(c.Context(), userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(user)
}

func (uc *UserController) Cancel() {
	// TODO
}
