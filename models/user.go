package models

import (
	"github.com/incrypt0/cokut-server/utils"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UID     string             `json:"uid,omitempty" bson:"uid,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty" `
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty" `
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	Address []string           `json:"address,omitempty" bson:"address,omitempty"`
	Admin   bool               `json:"admin,omitempty" bson:"admin,omitempty"`
}

// Prints Model Data in String
func (u *User) GetModelData() string {
	return utils.PrintModel(u)
}

// Real Validation
func (u *User) Validate() error {
	if (u.Name == "") || (len(u.Phone) < 10) || u.UID == "" {
		return errors.New("Not Validated")
	}
	return nil
}

// Basic Validate
func (u *User) ValidateBasic() error {
	if (u.Name == "") || (len(u.Phone) < 10) {
		return errors.New("Enter Valid Details")
	}
	return nil
}

// Basic Validate
func (u *User) ValidateEmail() error {
	if (u.Email == "") || (len(u.Email) < 5) {
		return errors.New("Enter A Valid Email")
	}
	return nil
}
