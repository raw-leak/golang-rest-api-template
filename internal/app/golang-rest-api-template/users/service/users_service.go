package service

import (
	"context"
	"errors"
	"log"

	"github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/dto"
	"github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/models"
)

// UserRepository represents the repository for users.
type UserRepository interface {
	InsertUser(ctx context.Context, user models.User) (string, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	// Add other repository methods as needed
}

type userService struct {
	userRepository UserRepository
}

// New creates a new UserService instance.
func New(userRepo UserRepository) *userService {
	return &userService{
		userRepository: userRepo,
	}
}

// CreateUser creates a new user.
func (us *userService) CreateUser(ctx context.Context, createUserDto dto.UserCreationDTO) (*models.User, error) {
	if createUserDto.Username == "" || createUserDto.Email == "" {
		return nil, errors.New("username and Email are required")
	}

	newUser := models.User{
		Username: createUserDto.Username,
		Email:    createUserDto.Email,
	}

	userId, err := us.userRepository.InsertUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	user, err := us.userRepository.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by their user ID.
func (us *userService) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	// Call the repository to retrieve the user by ID
	user, err := us.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Check if the user was found
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (us *userService) Cancel() {
	log.Println("Canceling user_service.")
}
