package models

import (
	"errors"

	"github.com/incrypt0/cokut-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RID            string             `json:"rid,omitempty" bson:"rid,omitempty"`
	UID            string             `json:"uid,omitempty" bson:"uid,omitempty"`
	Address        string             `json:"address,omitempty" bson:"address,omitempty"`
	Meals          []string           `json:"meals,omitempty" bson:"meals,omitempty"`
	Time           primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
	Price          float32            `json:"price,omitempty" bson:"price,omitempty"`
	DeliveryCharge float32            `json:"delivery_charge,omitempty" bson:"delivery_charge,omitempty"`
}

func (o *Order) GetModelData() string {
	return utils.ModelToString(o)
}

// Validate meal
func (o *Order) Validate() error {
	if (o.Address == "") || o.UID == "" || o.RID == "" {
		return errors.New("NOT_VALIDATED")
	}

	if len(o.Meals) == 0 {
		return errors.New("ITEMS_EMPTY")
	}

	return nil
}
