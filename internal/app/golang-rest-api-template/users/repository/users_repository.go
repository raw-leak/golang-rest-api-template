package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/MykhayloGusak/golang-rest-api-template/internal/app/golang-rest-api-template/users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	users *mongo.Collection
}

// NewUserRepository creates a new UserRepository instance.
func New(users *mongo.Collection) *userRepository {
	return &userRepository{
		users: users,
	}
}

// InsertUser inserts a new user into the repository.
func (ur *userRepository) InsertUser(ctx context.Context, user models.User) (string, error) {

	user.CreateAt = time.Now()
	user.UpdatedAt = time.Now()

	res, err := ur.users.InsertOne(ctx, user)
	if err != nil {
		log.Println("failed to insert user:", err)
		return "", err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("failed to assert the inserted ID to ObjectID")
		return "", errors.New("failed to assert the inserted ID to ObjectID")
	}

	return oid.Hex(), nil
}

// GetUserByID retrieves a user by their user ID from the repository.
func (ur *userRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	filter := bson.M{"_id": userID}

	err := ur.users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User not found
			return nil, nil
		}
		log.Println("Failed to retrieve user:", err)
		return nil, err
	}

	return &user, nil
}
