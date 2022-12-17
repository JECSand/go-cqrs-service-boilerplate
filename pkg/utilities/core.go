package utilities

import (
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

// StringToInt converts a string into an int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// CheckID validates a string id
func CheckID(idStr string) error {
	if idStr != "" && idStr != "000000000000000000000000" {
		return nil
	}
	return errors.New("invalid id string")
}

// NewID returns a new database id
func NewID() (uuid.UUID, error) {
	newId := primitive.NewObjectID()
	newIdStr := "00000000" + newId.Hex()
	return uuid.FromString(newIdStr)
}

// ToObjectID converts an uuid into an ObjectID
func ToObjectID(id uuid.UUID) (primitive.ObjectID, error) {
	idStr := strings.Replace(id.String(), "00000000", "", 1)
	if err := CheckID(idStr); err != nil {
		return primitive.ObjectID{}, err
	}
	return primitive.ObjectIDFromHex(idStr)
}
