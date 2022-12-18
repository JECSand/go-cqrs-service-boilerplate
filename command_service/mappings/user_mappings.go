package mappings

import (
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/entities"
	commandService "github.com/JECSand/go-cqrs-service-boilerplate/command_service/protos/user_command"
	kafkaMessages "github.com/JECSand/go-cqrs-service-boilerplate/protos/kafka"
	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToGrpcMessage(user *entities.User) *kafkaMessages.User {
	return &kafkaMessages.User{
		ID:        user.ID.String(),
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Root:      user.Root,
		Active:    user.Active,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func UserFromGrpcMessage(user *kafkaMessages.User) (*entities.User, error) {
	id, err := uuid.FromString(user.GetID())
	if err != nil {
		return nil, err
	}
	return &entities.User{
		ID:        id,
		Email:     user.GetEmail(),
		Username:  user.GetUsername(),
		Password:  user.GetPassword(),
		Root:      user.GetRoot(),
		Active:    user.GetActive(),
		CreatedAt: user.GetCreatedAt().AsTime(),
		UpdatedAt: user.GetUpdatedAt().AsTime(),
	}, nil
}

func CommandUserToGrpc(user *entities.User) *commandService.User {
	return &commandService.User{
		ID:        user.ID.String(),
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Root:      user.Root,
		Active:    user.Active,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
