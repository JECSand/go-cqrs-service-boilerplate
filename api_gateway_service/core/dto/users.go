package dto

import (
	"github.com/gofrs/uuid"
	"time"
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

// UserResponse ...
type UserResponse struct {
	ProductID   string    `json:"productId"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

func UserResponseFromGrpc(product *readerService.Product) *UserResponse {
	return &UserResponse{
		ProductID:   product.GetProductID(),
		Name:        product.GetName(),
		Description: product.GetDescription(),
		Price:       product.GetPrice(),
		CreatedAt:   product.GetCreatedAt().AsTime(),
		UpdatedAt:   product.GetUpdatedAt().AsTime(),
	}
}

// UsersListResponse ...
type UsersListResponse struct {
	TotalCount int64           `json:"totalCount" bson:"totalCount"`
	TotalPages int64           `json:"totalPages" bson:"totalPages"`
	Page       int64           `json:"page" bson:"page"`
	Size       int64           `json:"size" bson:"size"`
	HasMore    bool            `json:"hasMore" bson:"hasMore"`
	Users      []*UserResponse `json:"users" bson:"users"`
}

func UsersListResponseFromGrpc(listResponse *readerService.SearchRes) *UsersListResponse {
	list := make([]*UserResponse, 0, len(listResponse.GetProducts()))
	for _, product := range listResponse.GetProducts() {
		list = append(list, UserResponseFromGrpc(product))
	}
	return &UsersListResponse{
		TotalCount: listResponse.GetTotalCount(),
		TotalPages: listResponse.GetTotalPages(),
		Page:       listResponse.GetPage(),
		Size:       listResponse.GetSize(),
		HasMore:    listResponse.GetHasMore(),
		Users:      list,
	}
}
