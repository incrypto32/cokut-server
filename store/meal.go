package store

import (
	"errors"
	"log"

	"github.com/incrypt0/cokut-server/brokers/myerrors"
	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertMeal Function to insert Meals into meals collection.
func (s *Store) InsertMeal(m *models.Meal) (id string, err error) {
	c := s.mc
	rc := s.rc

	// Basic Validation
	if err = m.Validate(); err != nil {
		return id, err
	}

	m.Available = utils.NewBool(true)

	rid, err := primitive.ObjectIDFromHex(m.RID)

	if err != nil {
		return id, err
	}

	r, err := s.w.FindOne(rc, models.Restaurant{ID: rid})

	if err != nil {
		if errors.Is(myerrors.ErrNIL, err) {
			log.Println("NIL ERROR")
			return "", errors.New("restaurant doesn't exist")
		}

		return "", err
	}

	if r == nil {
		return "", errors.New("restaurant doesn't exist")
	}

	return s.w.Add(c, m)
}

// InsertSpecial Make a meal special
func (s *Store) InsertSpecial(id string, value bool) (string, error) {
	c := s.mc
	pid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return "", err
	}

	filter := models.Meal{ID: pid}
	update := models.Meal{Special: utils.NewBool(value)}

	r, err := s.w.FindOneAndUpdate(c, filter, update)

	if err != nil {
		log.Println(err)

		return "", err
	}

	a := (r.(*models.Meal))

	return a.ID.Hex(), err
}

//
func (s *Store) SearchMeal(searchText string) (l []interface{}, err error) {
	return s.w.Search(s.mc, models.Meal{}, searchText)
}

// GetMealsByRestaurant .
func (s *Store) GetMealsByRestaurant(rid string) (l []interface{}, err error) {
	return s.w.Get(s.mc, models.Meal{RID: rid})
}

// GetSpecialMeals .
func (s *Store) GetSpecialMeals() (l []interface{}, err error) {
	return s.w.Get(s.mc, models.Meal{Special: utils.NewBool(true)})
}

// GetSpiceyMeals .
func (s *Store) GetSpiceyMeals() (l []interface{}, err error) {
	return s.w.Get(s.mc, models.Meal{Spice: utils.NewBool(true)})
}

// deleteMeal
func (s *Store) DeleteMeal(id string) (int64, error) {
	return s.w.DeleteOne(s.mc, models.Meal{})
}
