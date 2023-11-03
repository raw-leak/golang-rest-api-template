package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/models"
	"github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/service"
	"github.com/MykhayloGusak/golang-rest-api-template/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {
	// Initialize the mock repository
	userRepositoryMock := mocks.UserRepository{}
	asserts := assert.New(t)

	// Setup expected values
	ctx := context.TODO()
	userDto := dto.UserCreationDTO{
		Username: "testuser",
		Email:    "test@example.com",
	}
	expectedUser := models.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: userDto.Username,
		Email:    userDto.Email,
	}
	expectedUserID := expectedUser.ID

	// Setup expectations
	userRepositoryMock.On("InsertUser", ctx, mock.AnythingOfType("models.User")).Return(expectedUserID, nil)
	userRepositoryMock.On("GetUserByID", ctx, expectedUserID).Return(&expectedUser, nil)

	// Call the service method
	userService := service.New(&userRepositoryMock)
	createdUser, err := userService.CreateUser(ctx, userDto)

	// Assert the expectations
	asserts.NoError(err)
	asserts.NotNil(createdUser)
	asserts.Equal(expectedUser.ID, createdUser.ID)
	asserts.Equal(expectedUser.Username, createdUser.Username)
	asserts.Equal(expectedUser.Email, createdUser.Email)

	// Assert that the expectations were met
	userRepositoryMock.AssertExpectations(t)
}

func TestCreateUser_Failure_MissingData(t *testing.T) {
	// Initialize the mock repository
	userRepositoryMock := mocks.UserRepository{}
	asserts := assert.New(t)

	// Setup expected values
	ctx := context.TODO()
	userDto := dto.UserCreationDTO{
		// Missing Username and Email
	}

	// Call the service method
	userService := service.New(&userRepositoryMock)
	createdUser, err := userService.CreateUser(ctx, userDto)

	// Assert there's an error and the user is nil
	asserts.Error(err)
	asserts.Nil(createdUser)
	asserts.EqualError(err, "username and Email are required")

	// Since we expect a failure before interacting with the repository,
	// no expectations are set on the mock, and hence, no need to assert expectations.
}

func TestCreateUser_UserAlreadyExists(t *testing.T) {
	// Initialize the mock repository
	userRepositoryMock := mocks.UserRepository{}
	asserts := assert.New(t)

	// Setup expected values
	ctx := context.TODO()
	userDto := dto.UserCreationDTO{
		Username: "testuser",
		Email:    "test@example.com",
	}

	// Setup expectations
	userRepositoryMock.On("InsertUser", ctx, mock.AnythingOfType("models.User")).Return("", errors.New("user already exists"))

	// Call the service method
	userService := service.New(&userRepositoryMock)
	createdUser, err := userService.CreateUser(ctx, userDto)

	// Assert the expectations
	asserts.Error(err)
	asserts.Nil(createdUser)
	asserts.EqualError(err, "user already exists")

	// Assert that the expectations were met
	userRepositoryMock.AssertExpectations(t)
}
