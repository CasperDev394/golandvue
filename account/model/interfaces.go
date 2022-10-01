package model

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
}

type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}
