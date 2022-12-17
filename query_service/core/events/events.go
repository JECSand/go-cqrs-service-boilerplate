package events

import (
	"github.com/gofrs/uuid"
	"time"
)

type UserEvents struct {
	CreateUser CreateUserEventHandler
	UpdateUser UpdateUserEventHandler
	DeleteUser DeleteUserEventHandler
}

func NewUserEvents(
	createUser CreateUserEventHandler,
	updateUser UpdateUserEventHandler,
	deleteUser DeleteUserEventHandler,
) *UserEvents {
	return &UserEvents{
		CreateUser: createUser,
		UpdateUser: updateUser,
		DeleteUser: deleteUser,
	}
}

type CreateUserEvent struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty" validate:"required,min=3,max=250"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=3,max=500"`
	Password  string    `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	Root      bool      `json:"root,omitempty" bson:"root,omitempty" validate:"required"`
	Active    bool      `json:"active,omitempty" bson:"active,omitempty" validate:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func NewCreateUserEvent(id string, email string, username string, password string, root bool, active bool, createdAt time.Time, updatedAt time.Time) *CreateUserEvent {
	return &CreateUserEvent{
		ID:        id,
		Email:     email,
		Username:  username,
		Password:  password,
		Root:      root,
		Active:    active,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type UpdateUserEvent struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty" validate:"required,min=3,max=250"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=3,max=500"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func NewUpdateUserEvent(id string, email string, username string, updatedAt time.Time) *UpdateUserEvent {
	return &UpdateUserEvent{
		ID:        id,
		Email:     email,
		Username:  username,
		UpdatedAt: updatedAt,
	}
}

type DeleteUserEvent struct {
	ID uuid.UUID `json:"id" bson:"_id,omitempty"`
}

func NewDeleteUserEvent(id uuid.UUID) *DeleteUserEvent {
	return &DeleteUserEvent{ID: id}
}
