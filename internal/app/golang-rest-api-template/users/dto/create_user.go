package dto

// UserCreationDTO represents the data required to create a new user.
type UserCreationDTO struct {
	Username string `json:"username"`
	Email     string `json:"email"`
	// Add other fields as needed
}
