package models

import (
	"github.com/incrypt0/cokut-server/utils"
	"github.com/pkg/errors"
)

// Address
type Address struct {
	Title        string `json:"title,omitempty" bson:"title,omitempty"`
	Zone         string `json:"zone,omitempty" bson:"zone,omitempty"`
	AddressLine1 string `json:"adl1,omitempty" bson:"adl1,omitempty"`
	AddressLine2 string `json:"adl2,omitempty" bson:"adl2,omitempty"`
	AddressLine3 string `json:"adl3,omitempty" bson:"adl3,omitempty"`
}

// User struct
type User struct {
	Admin     bool               `json:"admin,omitempty" bson:"admin,omitempty"`
	ID        string             `json:"id,omitempty" bson:"_id,omitempty"`
	UID       string             `json:"uid,omitempty" bson:"uid,omitempty"`
	GID       string             `json:"gid,omitempty" bson:"gid,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty" `
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty" `
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Address   []Address          `json:"address,omitempty" bson:"address,omitempty"`
	Addresses map[string]Address `json:"addresses,omitempty" bson:"addresses,omitempty"`
}

//GetModelData Prints Model Data in String
func (u *User) GetModelData() string {
	return utils.ModelToString(u)
}

//Validate Real Validation
func (u *User) Validate() error {
	if (u.Name == "") || (len(u.Phone) < 10) || u.UID == "" {
		return errors.New("Not Validated")
	}

	return nil
}

//ValidateBasic Basic Validate
func (u *User) ValidateBasic() error {
	if u.Name == "" {
		return errors.New("Enter Valid Details")
	}

	if u.Phone != "" && len(u.Phone) < 10 {
		return errors.New("Enter Valid Phone Number")
	}

	return nil
}

//ValidateEmail Basic Validate
func (u *User) ValidateEmail() error {
	if (u.Email == "") || (len(u.Email) < 5) {
		return errors.New("Enter A Valid Email")
	}

	return nil
}
