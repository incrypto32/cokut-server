package store

import "github.com/incrypt0/cokut-server/models"

// Function to insert users into userCollection
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

func (s *Store) GetAllRestaurants() (l []interface{}, err error) {
	return s.w.Get(s.rc, models.Restaurant{})
}

func (s *Store) GetAllHomeMade() (l []interface{}, err error) {
	return s.w.Get(s.rc, models.Restaurant{Type: "home"})
}
