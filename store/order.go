package store

import (
	"log"
	"math"
	"time"

	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrder creates a new order
func (s *Store) CreateOrder(o *models.Order, calculate bool) (po *models.Order, err error) {
	c := s.orders // Basic Validation

	o.Time = primitive.NewDateTimeFromTime(time.Now())

	if err = o.Validate(); err != nil {
		log.Println(err)

		return nil, err
	}

	if err = s.processOrder(o); err != nil {
		log.Println(err)

		return nil, err
	}

	if calculate {
		return o, err
	}

	o.StatusCode = s.orderCodes.Placed

	if id, err := s.w.Add(c, o); err != nil {
		return nil, err
	} else if o.ID, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}

	u, err := s.w.FindOne(s.uc, models.User{UID: o.UID})

	_ = s.bot.SendMessage(o.ToString(u.(*models.User)))

	return o, err
}

func (s *Store) processOrder(o *models.Order) error {
	mealCollection := s.mc
	ids := make([]string, 0, len(o.Items))

	for key := range o.Items {
		ids = append(ids, key)
	}

	rid, err := primitive.ObjectIDFromHex(o.RID)
	if err != nil {
		return err
	}

	result, err := s.w.FindOne(s.rc, models.Restaurant{ID: rid})
	r := result.(*models.Restaurant)

	if err != nil {
		return err
	} else if *(r.Closed) {
		return s.myerrors.ErrRestaurantClosed
	}

	l, err := s.w.GetMultipleByID(mealCollection, models.Meal{}, ids)

	if err != nil {
		return err
	}

	o.Restaurant = r

	if err = s.calculateDeliveryCharge(o, r.Location.Latitude, r.Location.Longitude); err != nil {
		return err
	}

	s.calculateOrderPrice(o, l)

	s.calculateServiceCharge(o)

	s.calculateTotal(o)

	return err
}

func (s *Store) calculateOrderPrice(o *models.Order, l []interface{}) {
	totalCount := 0

	for _, item := range l {
		meal := item.(*models.Meal)

		count := o.Items[meal.ID.Hex()]
		totalCount += count

		price := meal.Price * float64(count)

		o.Summary = append(o.Summary, &models.Summary{Meal: *meal, Count: count, Price: price})
		o.Price += price
	}

	o.TotalCount = totalCount
}

func (s *Store) calculateDeliveryCharge(o *models.Order, lat float64, long float64) error {
	dist := utils.Distance(o.Address.PlaceInfo.Latitude, o.Address.PlaceInfo.Longitude, lat, long)

	if dist > 20000 {
		return s.myerrors.ErrNotDeliverableArea
	}

	if dist <= 5000 {
		o.DeliveryCharge = 30.0

		return nil
	}

	o.DeliveryCharge = math.Round((dist / 1000) * 7)

	return nil
}

func (s *Store) calculateServiceCharge(o *models.Order) {
	if o.TotalCount <= 3 {
		o.ServiceCharge = 0
	} else {
		o.ServiceCharge = float64(o.TotalCount * 5)
	}
}

func (s *Store) calculateTotal(o *models.Order) {
	o.Total = o.Price + o.DeliveryCharge + o.ServiceCharge
}

// GetOrdersByUser user orders are returned
func (s *Store) GetOrdersByUser(uid string) (l []models.Order, err error) {
	return s.w.GetOrdersByUser(s.orders, 10, 1, uid)
}

// GetOrdersByUser user orders are returned
func (s *Store) GetPaginatedOrders(limit int, page int) (l []models.Order, err error) {
	return s.w.PaginatedOrders(s.orders, limit, page)
}

func (s *Store) ChangeOrderStatus(id string, statusCode int) (l *models.Order, err error) {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	o, err := s.w.FindOneAndUpdate(s.orders, models.Order{ID: pid}, models.Order{StatusCode: statusCode})

	if err != nil {
		log.Println(err)

		return nil, err
	}

	return o.(*models.Order), err
}
