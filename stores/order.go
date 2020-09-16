package stores

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
	collection *mongo.Collection
}

func GetOrderStore(collection *mongo.Collection) *OrderStore {
	return &OrderStore{
		collection: collection,
	}
}

// Function to insert Meals into meals collection
func (os *OrderStore) InsertOrder(o *models.Order, uid string) (id string, err error) {

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

	r := services.C.RestaurantsCollection.FindOne(context.Background(), bson.D{
		{Key: "_id", Value: rid},
	})

	if err = r.Err(); err != nil {
		return id, err
	}

	return services.Add(c, o)
}

func (os *OrderStore) GetOrders() (l []interface{}, err error) {

	return services.GetAll(os.collection, models.Order{})
}

func (os *OrderStore) GetUserOrders(uid string) (l []interface{}, err error) {
	return services.GetAll(os.collection, models.Order{UID: uid})
}
