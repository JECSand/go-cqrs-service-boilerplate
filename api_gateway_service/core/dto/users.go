package dto

import (
	"github.com/gofrs/uuid"
)

type CreateUserDTO struct {
	Email    string `json:"email" validate:"required,gte=0,lte=255"`
	Username string `json:"username" validate:"required,gte=0,lte=255"`
	Password string `json:"password" validate:"required,gte=0,lte=5000"`
}

type CreateUserResponseDTO struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type UpdateUserDTO struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Email    string    `json:"email" validate:"required,gte=0,lte=255"`
	Username string    `json:"username" validate:"required,gte=0,lte=255"`
	Password string    `json:"password" validate:"required,gte=0,lte=5000"`
}
