package store

import (
	"log"
	"time"

	"github.com/incrypt0/cokut-server/brokers/myerrors"
	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrder creates a new order
func (s *Store) CreateOrder(o *models.Order) (po *models.Order, err error) {
	c := s.orders // Basic Validation

	log.Println("2")

	o.Time = primitive.NewDateTimeFromTime(time.Now())

	if err = o.Validate(); err != nil {
		log.Println(err)

		return nil, err
	}

	if err = s.processOrder(o); err != nil {
		log.Println(err)

		return nil, err
	}

	if id, err := s.w.Add(c, o); err != nil {
		return nil, err
	} else if o.ID, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}

	return o, err
}

func (s *Store) processOrder(o *models.Order) error {
	mealCollection := s.mc
	ids := make([]string, 0, len(o.Items))

	log.Println("3")

	for key := range o.Items {
		ids = append(ids, key)
	}

	rid, err := primitive.ObjectIDFromHex(o.RID)
	if err != nil {
		return err
	}

	log.Println("4")

	result, err := s.w.FindOne(s.rc, models.Restaurant{ID: rid})
	r := result.(*models.Restaurant)

	if err != nil {
		return err
	} else if *(r.Closed) {
		return myerrors.ErrRestaurantClosed
	}

	l, err := s.w.GetMultipleByID(mealCollection, models.Meal{}, ids)

	s.calculateOrderPrice(o, l)

	s.calculateDeliveryCharge(o, r.Location.Latitude, r.Location.Longitude)

	s.calculateTotal(o)

	if err != nil {
		return err
	}

	return err
}

func (s *Store) calculateOrderPrice(o *models.Order, l []interface{}) {
	for _, item := range l {
		meal := item.(*models.Meal)

		count := o.Items[meal.ID.Hex()]

		price := meal.Price * float64(count)

		o.Summary = append(o.Summary, &models.Summary{Meal: *meal, Count: count, Price: price})
		o.Price += price
	}
}

func (s *Store) calculateDeliveryCharge(o *models.Order, lat float64, long float64) {
	dist := utils.Distance(o.Address.PlaceInfo.Latitude, o.Address.PlaceInfo.Longitude, lat, long)

	if dist <= 5000 {
		o.DeliveryCharge = 30.0

		return
	}

	o.DeliveryCharge = (dist / 1000) * 7
}

func (s *Store) calculateTotal(o *models.Order) {
	o.Total = o.Price + o.DeliveryCharge
}

// GetAllOrders Admin only function
func (s *Store) GetAllOrders() (l []interface{}, err error) {
	return s.w.Get(s.orders, models.Order{})
}

// GetOrdersByUser user orders are returned
func (s *Store) GetOrdersByUser(uid string) (l []interface{}, err error) {
	return s.w.Get(s.orders, bson.M{"uid": uid})
}

// GetOrdersByUser user orders are returned
func (s *Store) GetPaginatedOrders(limit int, page int) (l []models.OrderSummary, err error) {
	return s.w.PaginatedOrders(s.orders, limit, page)
}
