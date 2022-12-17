package commands

import (
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/dto"
	"github.com/gofrs/uuid"
)

type UserCommands struct {
	CreateProduct CreateUserCmdHandler
	UpdateProduct UpdateUserCmdHandler
	DeleteProduct DeleteUserCmdHandler
}

func NewProductCommands(create CreateUserCmdHandler, update UpdateUserCmdHandler, delete DeleteUserCmdHandler) *UserCommands {
	return &UserCommands{
		CreateProduct: create,
		UpdateProduct: update,
		DeleteProduct: delete,
	}
}

// CreateUserCommand ...
type CreateUserCommand struct {
	CreateDto *dto.CreateUserDTO
}

func NewCreateUserCommand(createDto *dto.CreateUserDTO) *CreateUserCommand {
	return &CreateUserCommand{CreateDto: createDto}
}

// UpdateUserCommand ...
type UpdateUserCommand struct {
	UpdateDto *dto.UpdateUserDTO
}

func NewUpdateUserCommand(updateDto *dto.UpdateUserDTO) *UpdateUserCommand {
	return &UpdateUserCommand{UpdateDto: updateDto}
}

// DeleteUserCommand ...
type DeleteUserCommand struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func NewDeleteProductCommand(userID uuid.UUID) *DeleteUserCommand {
	return &DeleteUserCommand{ID: userID}
}
