package models

import (
	"errors"
	"fmt"

	"github.com/incrypt0/cokut-server/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	return services.PrintModel(m)
}

// Validate meal
func (m *Meal) Validate() error {

	if m.Name == "" || (m.Price <= 0) || m.DisplayPrice <= 0 || m.RestID == "" {
		return errors.New("Not Validated")
	}

	return nil
}

// Function to insert Meals into meals collection
func InsertMeal(m *Meal) (id string, err error) {

	//  Getting the user colection
	var c *mongo.Collection = services.C.MealsCollection

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

	return services.Add(c, m)
}

// Insert a special item
func InsertSpecial(id string) (result_id primitive.ObjectID, err error) {
	var c *mongo.Collection = services.C.MealsCollection
	rest := new(Restaurant)
	mongoid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		err = errors.New("An error occured")
		return result_id, err
	}
	r := c.FindOneAndUpdate(ctx, bson.D{
		{Key: "_id", Value: mongoid},
	}, bson.M{"$set": bson.D{
		{Key: "special", Value: true},
	}})

	if err = r.Err(); err != nil {
		fmt.Println(err)
		err = errors.New("An error occured")
		return result_id, err
	}

	if err = r.Decode(rest); err != nil {
		fmt.Println(err)

		err = errors.New("An error occured")
		return result_id, err
	}
	result_id = rest.ID
	return result_id, err
}

func GetMeals(rid string) (l []interface{}, err error) {
	return services.GetAll(services.C.MealsCollection, Meal{RestID: rid})
}

func GetSpecials() (l []interface{}, err error) {
	fmt.Println("Test 1")
	return services.GetAll(services.C.MealsCollection, Meal{Special: true})
}

func GetSpicey() (l []interface{}, err error) {
	fmt.Println("Test Spicey")
	return services.GetAll(services.C.MealsCollection, Meal{Spicey: true})
}
