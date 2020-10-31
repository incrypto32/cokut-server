package models

import (
	"errors"

	"github.com/incrypt0/cokut-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meal struct {
	IsVeg        *bool              `json:"isVeg,omitempty" bson:"isVeg,omitempty"`
	Special      *bool              `json:"special,omitempty" bson:"special,omitempty"`
	Spice        *bool              `json:"spice,omitempty" bson:"spice,omitempty"`
	Available    *bool              `json:"available,omitempty" bson:"available,omitempty"`
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	RID          string             `json:"rid,omitempty" bson:"rid,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty" `
	Price        float64            `json:"price,omitempty" bson:"price,omitempty" `
	StrikedPrice float64            `json:"strikedPrice,omitempty" bson:"strikedPrice,omitempty"`
	Keywords     []string           `json:"keywords,omitempty" bson:"keywords,omitempty"`
}

func (m *Meal) GetModelData() string {
	return utils.ModelToString(m)
}

// Validate meal .
func (m *Meal) Validate() error {

	if m.Name == "" || (m.Price <= 0) || m.RID == "" {
		return errors.New("NOT_VALIDATED")
	}

	return nil
}
