package models

import (
	"github.com/google/uuid"
	"time"
)

// User is the data type for user object
type User struct {
	ID         uuid.UUID `json:"id" sql:"id"`
	Email      string    `json:"email" validate:"required" sql:"email"`
	Password   string    `json:"password" validate:"required" sql:"password"`
	Username   string    `json:"username" sql:"username"`
	TokenHash  string    `json:"tokenhash" sql:"tokenhash"`
	IsVerified bool      `json:"isverified" sql:"isverified"`
	CreatedAt  time.Time `json:"createdat" sql:"createdat"`
	UpdatedAt  time.Time `json:"updatedat" sql:"updatedat"`
}

type UpdateResult struct {
	ID       uuid.UUID `json:"id" sql:"id"`
	Username string    `json:"username" sql:"username"`
	Email    string    `json:"email" validate:"required" sql:"email"`
	Password string    `json:"password" validate:"required" sql:"password"`
}
