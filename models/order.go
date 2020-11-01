package models

import (
	"github.com/incrypt0/cokut-server/brokers/myerrors"
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
	Address        *Address           `json:"address,omitempty" bson:"address,omitempty"`
	Items          map[string]int     `json:"items,omitempty" bson:"items,omitempty"`
	Summary        []*Summary         `json:"summary,omitempty" bson:"summary,omitempty"`
	User           *User              `json:"user,omitempty" bson:"user,omitempty"`
	Restaurant     *Restaurant        `json:"restaurant,omitempty" bson:"restaurant,omitempty"`
	Time           primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
	Price          float64            `json:"price,omitempty" bson:"price,omitempty"`
	Total          float64            `json:"total,omitempty" bson:"total,omitempty"`
	DeliveryCharge float64            `json:"deliveryCharge,omitempty" bson:"deliveryCharge,omitempty"`
	ServiceCharge  float64            `json:"serviceCharge,omitempty" bson:"serviceCharge,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
	StatusCode     int                `json:"statusCode,omitempty" bson:"statusCode,omitempty"`
	TotalCount     int                `json:"totalCount,omitempty" bson:"totalCount,omitempty"`
}

func (o *Order) GetModelData() string {
	return utils.ModelToString(o)
}

// Validate meal
func (o *Order) Validate() error {
	if o.UID == "" || o.RID == "" || o.Items == nil || o.Address == nil {
		return myerrors.ErrOrderNotValidated
	}

	return nil
}

// Validate meal
func (o *Order) ValidateBasic() error {
	if o.UID == "" || o.RID == "" || o.Items == nil {
		return myerrors.ErrOrderNotValidated
	}

	return nil
}
