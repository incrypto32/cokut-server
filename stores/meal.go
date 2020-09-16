package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MealStore struct {
	collection  *mongo.Collection
	rcollection *mongo.Collection
}

func NewMealStore(collection *mongo.Collection, rcollection *mongo.Collection) *MealStore {
	return &MealStore{
		collection:  collection,
		rcollection: rcollection,
	}
}

// Function to insert Meals into meals collection
func (ms *MealStore) Insert(m *models.Meal) (id string, err error) {

	//  Getting the user colection
	var c *mongo.Collection = ms.collection

	// Basic Validation
	if err = m.Validate(); err != nil {
		return id, err
	}

	rid, err := primitive.ObjectIDFromHex(m.RestID)
	if err != nil {
		return id, err
	}

	r := ms.rcollection.FindOne(context.Background(), bson.D{
		{Key: "_id", Value: rid},
	})

	if err = r.Err(); err != nil {
		return id, err
	}

	return services.Add(c, m)
}

// Insert a special item
func (ms *MealStore) InsertSpecial(id string) (result_id string, err error) {
	var c *mongo.Collection = ms.collection
	rest := new(models.Restaurant)
	mongoid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		err = errors.New("An error occured")
		return result_id, err
	}
	r := c.FindOneAndUpdate(context.TODO(), bson.D{
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
	result_id = rest.ID.Hex()
	return result_id, err
}

func (ms *MealStore) GetByRestaurant(rid string) (l []interface{}, err error) {
	return services.GetAll(ms.collection, models.Meal{RestID: rid})
}

func (ms *MealStore) GetSpecials() (l []interface{}, err error) {
	fmt.Println("Test 1")
	return services.GetAll(ms.collection, models.Meal{Special: true})
}

func (ms *MealStore) GetSpicey() (l []interface{}, err error) {
	fmt.Println("Test Spicey")
	return services.GetAll(ms.collection, models.Meal{Spicey: true})
}
