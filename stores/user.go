package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	collection *mongo.Collection
}

func NewUserStore(collection *mongo.Collection) *UserStore {
	return &UserStore{
		collection: collection,
	}
}

// Function to insert users into userCollection
func (us *UserStore) Insert(u *models.User) (id string, err error) {
	//  Getting the user colection
	c := us.collection

	// Basic Validation
	if err = u.ValidateBasic(); err != nil {
		return id, err
	}

	// Check if email is null
	if u.Email != "" {
		e := c.FindOne(context.Background(), bson.D{
			{Key: "email", Value: u.Email},
		})

		if e.Err() != mongo.ErrNoDocuments {
			err = errors.New("Email Address is already associated with another account")
			return id, err
		}
	}

	// Check if phone is null
	if u.Phone != "" {
		r := c.FindOne(context.Background(), bson.D{
			{Key: "phone", Value: u.Phone},
		})

		if r.Err() != mongo.ErrNoDocuments {
			err = errors.New("Phone number is already associated with another account")
			return id, err
		}
	}

	if err != nil {
		return id, err
	}

	return services.Add(c, u)
}

// Check User existence
func (us *UserStore) CheckUserExistence(phone string) bool {
	var val bool = false
	fmt.Println("CheckUser called with phone : ", phone)
	c := us.collection
	r := c.FindOne(context.Background(), bson.D{
		{Key: "phone", Value: phone},
	})

	if r.Err() != mongo.ErrNoDocuments {
		val = true
	}
	fmt.Println(phone, " exists : ", val)
	return val

}
