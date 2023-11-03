package models

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Username  string    `bson:"username"`
	Email     string    `bson:"email"`
	CreateAt  time.Time `bson:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
	// Add other user fields as needed
}
