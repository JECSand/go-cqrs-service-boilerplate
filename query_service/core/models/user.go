package models

import (
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/utilities"
	queryService "github.com/JECSand/go-cqrs-service-boilerplate/query_service/protos/user_query"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty" validate:"required,min=3,max=250"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=3,max=500"`
	Password  string    `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	Root      bool      `json:"root,omitempty" bson:"root,omitempty" validate:"required"`
	Active    bool      `json:"active,omitempty" bson:"active,omitempty" validate:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// UsersList products list response with pagination
type UsersList struct {
	TotalCount int64   `json:"totalCount" bson:"totalCount"`
	TotalPages int64   `json:"totalPages" bson:"totalPages"`
	Page       int64   `json:"page" bson:"page"`
	Size       int64   `json:"size" bson:"size"`
	HasMore    bool    `json:"hasMore" bson:"hasMore"`
	Users      []*User `json:"users" bson:"users"`
}

func NewUserListWithPagination(users []*User, count int64, pagination *utilities.Pagination) *UsersList {
	return &UsersList{
		TotalCount: count,
		TotalPages: int64(pagination.GetTotalPages(int(count))),
		Page:       int64(pagination.GetPage()),
		Size:       int64(pagination.GetSize()),
		HasMore:    pagination.GetHasMore(int(count)),
		Users:      users,
	}
}

func UserToGrpcMessage(user *User) *queryService.User {
	return &queryService.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Root:      user.Root,
		Active:    user.Active,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func UserListToGrpc(users *UsersList) *queryService.SearchRes {
	list := make([]*queryService.User, 0, len(users.Users))
	for _, user := range users.Users {
		list = append(list, UserToGrpcMessage(user))
	}
	return &queryService.SearchRes{
		TotalCount: users.TotalCount,
		TotalPages: users.TotalPages,
		Page:       users.Page,
		Size:       users.Size,
		HasMore:    users.HasMore,
		Users:      list,
	}
}
