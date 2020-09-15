package models

import (
	"errors"
	"fmt"

	"github.com/incrypt0/cokut-server/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meal struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RestID       string             `json:"rid,omitempty" bson:"rid,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty" `
	Price        float32            `json:"price,omitempty" bson:"price,omitempty" `
	DisplayPrice float32            `json:"display_price,omitempty" bson:"display_price,omitempty"`
	IsVeg        bool               `json:"isVeg,omitempty" bson:"isVegemail,omitempty"`
}

func (m *Meal) GetModelData() string {
	return services.PrintModel(m)
}

// Validate meal
func (m *Meal) Validate() error {
	fmt.Println(services.PrintModel(m))
	fmt.Println(m.Price)
	if m.Name == "" || (m.Price <= 0) || m.DisplayPrice <= 0 || m.RestID == "" {
		return errors.New("Not Validated")
	}

	return nil
}

// Function to insert Meals into meals collection
func InsertMeal(m *Meal) (id primitive.ObjectID, err error) {

	//  Getting the user colection
	c := services.C.MealsCollection

	// Basic Validation
	if err = m.Validate(); err != nil {
		return id, err
	}

	rid, err := primitive.ObjectIDFromHex(m.RestID)
	if err != nil {
		return id, err
	}

	r := services.C.RestaurantsCollection.FindOne(ctx, bson.D{
		{Key: "_id", Value: rid},
	})

	if err = r.Err(); err != nil {
		return id, err
	}

	fmt.Println("_____Validated_____")
	fmt.Println(services.PrintModel(m))
	return services.Add(c, m)
}
