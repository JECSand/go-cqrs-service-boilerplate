package data

import (
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/utilities"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type entity interface {
	getID() primitive.ObjectID
}

// userEntity structures a user BSON document to save in a users collection
type userEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username,omitempty" validate:"required,min=3,max=500"`
	Password  string             `bson:"password,omitempty" validate:"required"`
	Email     string             `bson:"email,omitempty" validate:"required,min=3,max=250"`
	Root      bool               `bson:"root,omitempty"`
	Active    bool               `bson:"active,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}

// getID returns the unique identifier of the userEntity
func (u *userEntity) getID() primitive.ObjectID {
	return u.ID
}

// newUserEntity initializes a new pointer to a userModel struct from a pointer to a JSON User struct
func newUserEntity(u *models.User) (um *userEntity, err error) {
	um = &userEntity{
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		Root:      u.Root,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	if utilities.CheckID(u.ID) == nil {
		um.ID, err = utilities.LoadObjectIDString(u.ID)
	}
	return
}

// toRoot creates and return a new pointer to a User JSON struct from a pointer to a BSON userModel
func (u *userEntity) toRoot() *models.User {
	um := &models.User{
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		Root:      u.Root,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	if utilities.CheckID(u.ID.Hex()) == nil {
		um.ID = utilities.LoadUUIDString(u.ID)
	}
	return um
}

func loadUserEntities(ms []*userEntity) (users []*models.User) {
	for _, m := range ms {
		users = append(users, m.toRoot())
	}
	return
}
