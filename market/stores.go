package market

import (
	"github.com/incrypt0/cokut-server/models"
)

type UserStore interface {
	CheckUserExistence(phone string) bool
	Insert(u *models.User) (id string, err error)
}

type RestaurantStore interface {
	Insert(r *models.Restaurant) (id string, err error)
	GetAll() (l []interface{}, err error)
	GetAllHomeMade() (l []interface{}, err error)
}

type OrderStore interface {
	Insert(o *models.Order, uid string) (id string, err error)
	GetAll() (l []interface{}, err error)
	GetByUser(uid string) (l []interface{}, err error)
}

type MealStore interface {
	Insert(m *models.Meal) (id string, err error)
	InsertSpecial(id string) (result_id string, err error)
	GetByRestaurant(rid string) (l []interface{}, err error)
	GetSpecials() (l []interface{}, err error)
	GetSpicey() (l []interface{}, err error)
}
