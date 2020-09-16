package stores

import (
	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantStore struct {
	collection *mongo.Collection
}

func GetRestaurantStore(collection *mongo.Collection) *RestaurantStore {
	return &RestaurantStore{
		collection: collection,
	}
}

// Function to insert users into userCollection
func (rs *RestaurantStore) InsertRestaurant(r *models.Restaurant) (id string, err error) {
	//  Getting the user colection
	c := rs.collection

	// Basic Validation
	if err = r.Validate(); err != nil {
		return id, err
	}

	if err != nil {
		return id, err
	}

	return services.Add(c, r)
}

func (rs *RestaurantStore) GetAllRestaurants() (l []interface{}, err error) {
	return services.GetAll(rs.collection, models.Restaurant{})
}

func (rs *RestaurantStore) GetHomeMade() (l []interface{}, err error) {
	return services.GetAll(rs.collection, models.Restaurant{Type: "home"})
}
