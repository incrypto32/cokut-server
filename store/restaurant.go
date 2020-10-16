package store

import (
	"github.com/incrypt0/cokut-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRestaurant Function to insert restaurants into db
func (s *Store) InsertRestaurant(r *models.Restaurant) (id string, err error) {
	//  Getting the user colection
	c := s.rc

	// Basic Validation
	if err = r.Validate(); err != nil {
		return id, err
	}

	if err != nil {
		return id, err
	}

	return s.w.Add(c, r)
}

// GetAllRestaurants Function to get all restaurants
func (s *Store) GetAllRestaurants() (l []interface{}, err error) {
	return s.w.Get(s.rc, models.Restaurant{})
}

// GetAllRestaurants Function to get all restaurants
func (s *Store) GetAllRegularRestaurants() (l []interface{}, err error) {
	return s.w.Get(s.rc, models.Restaurant{Type: "regular"})
}

// GetAllHomeMade Function to get all home food centers
func (s *Store) GetAllHomeMade() (l []interface{}, err error) {
	return s.w.Get(s.rc, models.Restaurant{Type: "home"})
}

// Update Restaurant

func (s *Store) UpdateRestaurant(id string, restaurant models.Restaurant) (l interface{}, err error) {
	pid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return s.w.FindOneAndUpdate(s.rc, models.Restaurant{ID: pid}, restaurant)
}
