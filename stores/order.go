package store

import (
	"context"
	"time"

	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderStore struct {
	collection  *mongo.Collection
	rcollection *mongo.Collection
}

func NewOrderStore(collection *mongo.Collection, rcollection *mongo.Collection) *OrderStore {
	return &OrderStore{
		collection:  collection,
		rcollection: rcollection,
	}
}

// Function to insert Meals into meals collection
func (os *OrderStore) Insert(o *models.Order, uid string) (id string, err error) {

	o.Time = primitive.Timestamp{T: uint32(time.Now().Unix())}
	o.UID = uid

	var c *mongo.Collection = os.collection

	// Basic Validation
	if err = o.Validate(); err != nil {
		return id, err
	}

	rid, err := primitive.ObjectIDFromHex(o.RID)
	if err != nil {
		return id, err
	}

	r := os.collection.FindOne(context.Background(), bson.D{
		{Key: "_id", Value: rid},
	})

	if err = r.Err(); err != nil {
		return id, err
	}

	return services.Add(c, o)
}

func (os *OrderStore) GetAll() (l []interface{}, err error) {
	return services.GetAll(os.collection, models.Order{})
}

func (os *OrderStore) GetByUser(uid string) (l []interface{}, err error) {
	return services.GetAll(os.collection, models.Order{UID: uid})
}
