package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/incrypt0/cokut-server/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name" `
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty" `
	LogoUrl string             `json:"logo,omitempty" bson:"logo,omitempty"`
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	Address string             `json:"address,omitempty" bson:"address,omitempty"`
	Type    string             `json:"type,omitempty" bson:"type,omitempty"`
}

func (r *Restaurant) GetModelData() string {
	return services.PrintModel(r)
}

func (r *Restaurant) Validate() error {
	fmt.Println(services.PrintModel(r))

	if r.Name == "" || (len(r.Phone) < 10) || r.Address == "" {
		return errors.New("Not Validated")
	}
	return nil
}

// Function to insert users into userCollection
func InsertRestaurant(r *Restaurant) (id primitive.ObjectID, err error) {
	//  Getting the user colection
	c := services.C.RestaurantsCollection

	// Basic Validation
	if err = r.Validate(); err != nil {
		return id, err
	}

	if err != nil {
		return id, err
	}

	fmt.Println(services.PrintModel(r))
	return services.Add(c, r)
}

func GetAllRestaurants() (l []Restaurant, err error) {

	cur, err := services.GetAll(services.C.RestaurantsCollection, Restaurant{})

	for cur.Next(ctx) {
		rest := new(Restaurant)

		if err = cur.Decode(rest); err != nil {
			log.Println(err)
			return l, errors.New("An Error Occured")
		}

		l = append(l, *rest)
	}

	defer cur.Close(ctx)
	return l, err
}

func GetAllRestaurants2() (l []interface{}, err error) {

	return services.GetAll2(services.C.RestaurantsCollection, Restaurant{})
}
