package models

import "time"

// Blacklist is a root struct that is used to store the json encoded data for/from a mongodb blacklist doc.
type Blacklist struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	AuthToken string    `json:"authToken,omitempty" bson:"authToken,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// GetID returns the unique identifier of the Blacklist
func (b *Blacklist) GetID() string {
	return b.ID
}
