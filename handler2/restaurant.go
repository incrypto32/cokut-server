package handler2

import (
	"log"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Add a single restaurant
func (h *Handler) addRestaurant(c echo.Context) (err error) {
	r := new(models.Restaurant)
	return h.Add(c, r, func(r models.Model) (string, error) {
		return h.store.InsertRestaurant(r.(*models.Restaurant))
	})
}

// Get all restaurants in the db
func (h *Handler) getAllRestaurants(c echo.Context) (err error) {

	l, err := h.store.GetAllRestaurants()

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	return c.JSON(http.StatusOK, l)
}

//get Home
func (h *Handler) getHomeMadeRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllHomeMade)
}
