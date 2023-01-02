package entities

import (
	"errors"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Root      bool      `json:"root,omitempty"`
	Active    bool      `json:"active,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// Authenticate compares an input password with the hashed password stored in the User model
func (u *User) Authenticate(checkPassword string) error {
	if len(u.Password) != 0 {
		password := []byte(u.Password)
		cPassword := []byte(checkPassword)
		return bcrypt.CompareHashAndPassword(password, cPassword)
	}
	return errors.New("user password is missing")
}

// HashPassword hashes a User Password
func (u *User) HashPassword() error {
	if len(u.Password) != 0 {
		password := []byte(u.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
		return nil
	}
	return errors.New("user password is missing")
}
