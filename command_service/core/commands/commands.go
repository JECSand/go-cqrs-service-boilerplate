package commands

import (
	"github.com/gofrs/uuid"
)

type UserCommands struct {
	CreateUser CreateUserCmdHandler
	UpdateUser UpdateUserCmdHandler
	DeleteUser DeleteUserCmdHandler
}

func NewUserCommands(createUser CreateUserCmdHandler, updateUser UpdateUserCmdHandler, deleteUser DeleteUserCmdHandler) *UserCommands {
	return &UserCommands{
		CreateUser: createUser,
		UpdateUser: updateUser,
		DeleteUser: deleteUser,
	}
}

type CreateUserCommand struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Email    string    `json:"email" validate:"required,gte=0,lte=255"`
	Username string    `json:"username" validate:"required,gte=0,lte=5000"`
	Password string    `json:"password" validate:"required"`
	Root     bool      `json:"root"`
	Active   bool      `json:"active"`
}

func NewCreateUserCommand(id uuid.UUID, email string, username string, password string, root bool, active bool) *CreateUserCommand {
	return &CreateUserCommand{
		ID:       id,
		Email:    email,
		Username: username,
		Password: password,
		Root:     root,
		Active:   active,
	}
}

type UpdateUserCommand struct {
	ID       uuid.UUID `json:"id" validate:"required,gte=0,lte=255"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func NewUpdateUserCommand(id uuid.UUID, email string, username string) *UpdateUserCommand {
	return &UpdateUserCommand{
		ID:       id,
		Email:    email,
		Username: username,
	}
}

type DeleteUserCommand struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func NewDeleteUserCommand(id uuid.UUID) *DeleteUserCommand {
	return &DeleteUserCommand{ID: id}
}
