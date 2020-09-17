package models

import (
	"errors"

	"github.com/incrypt0/cokut-server/workers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meal struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RestID       string             `json:"rid,omitempty" bson:"rid,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty" `
	Price        float32            `json:"price,omitempty" bson:"price,omitempty" `
	DisplayPrice float32            `json:"display_price,omitempty" bson:"display_price,omitempty"`
	IsVeg        bool               `json:"isVeg,omitempty" bson:"isVeg,omitempty"`
	Special      bool               `json:"special,omitempty" bson:"special,omitempty"`
	Spicey       bool               `json:"spicey,omitempty" bson:"spicey,omitempty"`
}

func (m *Meal) GetModelData() string {
	return workers.PrintModel(m)
}

// Validate meal
func (m *Meal) Validate() error {

	if m.Name == "" || (m.Price <= 0) || m.DisplayPrice <= 0 || m.RestID == "" {
		return errors.New("Not Validated")
	}

	return nil
}
