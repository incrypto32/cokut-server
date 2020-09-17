package stores

import (
	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/workers"

	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantStore struct {
	collection *mongo.Collection
}

func NewRestaurantStore(collection *mongo.Collection) *RestaurantStore {
	return &RestaurantStore{
		collection: collection,
	}
}

// Function to insert users into userCollection
func (rs *RestaurantStore) Insert(r *models.Restaurant) (id string, err error) {

	//  Getting the user colection
	c := rs.collection

	// Basic Validation
	if err = r.Validate(); err != nil {
		return id, err
	}

	if err != nil {
		return id, err
	}

	return workers.Add(c, r)
}

func (rs *RestaurantStore) GetAll() (l []interface{}, err error) {
	return workers.Get(rs.collection, models.Restaurant{})
}

func (rs *RestaurantStore) GetAllHomeMade() (l []interface{}, err error) {
	return workers.Get(rs.collection, models.Restaurant{Type: "home"})
}
