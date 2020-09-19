package store

import (
	"time"

	"github.com/incrypt0/cokut-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrder creates a new order
func (s *Store) CreateOrder(o *models.Order) (id string, err error) {
	c := s.orders // Basic Validation

	o.Time = primitive.NewDateTimeFromTime(time.Now())

	if err = o.Validate(); err != nil {
		return id, err
	}

	return s.w.Add(c, o)
}

// GetAllOrders Admin only function
func (s *Store) GetAllOrders() (l []interface{}, err error) {
	return s.w.Get(s.orders, models.Order{})
}

// GetOrdersByUser user orders are returned
func (s *Store) GetOrdersByUser(uid string) (l []interface{}, err error) {
	return s.w.Get(s.orders, models.Order{UID: uid})
}
