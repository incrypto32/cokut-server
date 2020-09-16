package handler

import (
	"fmt"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Add a single restaurant
func (h *Handler) addRestaurant(c echo.Context) (err error) {
	r := new(models.Restaurant)
	return h.Add(c, r, func(r models.Model) (string, error) {
		return models.InsertRestaurant(r.(*models.Restaurant))
	})
}

// Get all restaurants in the db
func (h *Handler) getAllRestaurants(c echo.Context) (err error) {

	l, err := models.GetAllRestaurants()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	return c.JSON(http.StatusOK, l)
}

//get Home
func (h *Handler) getHomeMadeRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, models.GetHomeMade)
}
