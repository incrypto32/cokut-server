package models

import (
	"errors"

	"github.com/incrypt0/cokut-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Closed   bool               `json:"closed,omitempty" bson:"closed,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty" `
	LogoURL  string             `json:"logo,omitempty" bson:"logo,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Address  string             `json:"address,omitempty" bson:"address,omitempty"`
	Type     string             `json:"type,omitempty" bson:"type,omitempty"`
	Keywords []string           `json:"keywords,omitempty" bson:"keywords,omitempty"`
}

func (r *Restaurant) GetModelData() string {
	return utils.ModelToString(r)
}

func (r *Restaurant) Validate() error {
	if r.Name == "" || (len(r.Phone) < 10) || r.Address == "" {
		return errors.New("NOT_VALIDATED")
	}

	return nil
}
