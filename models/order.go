package models

import (
	"errors"

	"github.com/incrypt0/cokut-server/workers"
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
	return workers.PrintModel(o)
}

// Validate meal
func (o *Order) Validate() error {

	if (o.Address == "") || o.UID == "" || o.RID == "" {
		return errors.New("Not Validated")
	}
	if len(o.Meals) <= 0 {
		return errors.New("Items Empty")
	}

	return nil
}

// // Function to insert Meals into meals collection
// func InsertOrder(o *Order, uid string) (id string, err error) {

// 	o.Time = primitive.Timestamp{T: uint32(time.Now().Unix())}
// 	o.UID = uid
// 	//  Getting the user colection
// 	var c *mongo.Collection = workers.C.OrderCollecton

// 	// Basic Validation
// 	if err = o.Validate(); err != nil {
// 		return id, err
// 	}

// 	rid, err := primitive.ObjectIDFromHex(o.RID)
// 	if err != nil {
// 		return id, err
// 	}

// 	r := workers.C.RestaurantsCollection.FindOne(ctx, bson.D{
// 		{Key: "_id", Value: rid},
// 	})

// 	if err = r.Err(); err != nil {
// 		return id, err
// 	}

// 	return workers.Add(c, o)
// }

// func GetOrders() (l []interface{}, err error) {

// 	return workers.GetAll(workers.C.OrderCollecton, Order{})
// }

// func GetUserOrders(uid string) (l []interface{}, err error) {
// 	return workers.GetAll(workers.C.OrderCollecton, Order{UID: uid})
// }
