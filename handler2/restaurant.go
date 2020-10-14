package handler2

import (
	"log"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Add a single restaurant
func (h *Handler) addRestaurant(c echo.Context) (err error) {
	r := new(models.Restaurant)

	return h.Add(c, r, func(r models.Model) (interface{}, error) {
		return h.store.InsertRestaurant(r.(*models.Restaurant))
	})
}
func (h *Handler) addRestaurantForm(c echo.Context) (err error) {
	// r := new(models.Restaurant)
	log.Println(c.FormParams())

	// return h.store.InsertRestaurant(r.(*models.Restaurant))
	return err
}

// Get all restaurants in the db
func (h *Handler) getAllRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllRestaurants)
}

// Get all restaurants in the db
func (h *Handler) getAllRegularRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllRegularRestaurants)
}

//get Home
func (h *Handler) getHomeMadeRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllHomeMade)
}
