package models

import (
	"errors"

	"github.com/incrypt0/cokut-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Summary struct {
	Meal  Meal    `json:"meal,omitempty" bson:"meal,omitempty"`
	Count int     `json:"count,omitempty" bson:"count,omitempty"`
	Price float64 `json:"price,omitempty" bson:"price,omitempty"`
}

type Order struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	RID            string             `json:"rid,omitempty" bson:"rid,omitempty"`
	UID            string             `json:"uid,omitempty" bson:"uid,omitempty"`
	Address        Address            `json:"address,omitempty" bson:"address,omitempty"`
	Items          map[string]int     `json:"items,omitempty" bson:"items,omitempty"`
	Summary        []Summary          `json:"summary,omitempty" bson:"summary,omitempty"`
	Time           primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
	Price          float64            `json:"price,omitempty" bson:"price,omitempty"`
	Total          float64            `json:"total,omitempty" bson:"total,omitempty"`
	DeliveryCharge float64            `json:"delivery_charge,omitempty" bson:"delivery_charge,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

func (o *Order) GetModelData() string {
	return utils.ModelToString(o)
}

// Validate meal
func (o *Order) Validate() error {
	if (o.Address.Zone == "") || o.UID == "" || o.RID == "" {
		return errors.New("NOT_VALIDATED")
	}

	return nil
}
