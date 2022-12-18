package queries

import (
	"github.com/gofrs/uuid"
)

type UserQueries struct {
	GetUserById GetUserByIdHandler
}

func NewUserQueries(getById GetUserByIdHandler) *UserQueries {
	return &UserQueries{GetUserById: getById}
}

type GetUserByIdQuery struct {
	ID uuid.UUID `json:"id" validate:"required,gte=0,lte=255"`
}

func NewGetUserByIdQuery(id uuid.UUID) *GetUserByIdQuery {
	return &GetUserByIdQuery{ID: id}
}
