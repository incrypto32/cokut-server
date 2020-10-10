package handler2

import (
	"log"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Add a meal to the db.
func (h *Handler) addMeal(c echo.Context) (err error) {
	r := new(models.Meal)

	return h.Add(c, r, func(r models.Model) (interface{}, error) {
		return h.store.InsertMeal(r.(*models.Meal))
	})
}

// Get all meals from the database with the given restaurant ID.
func (h *Handler) getMeals(c echo.Context) (err error) {
	return h.getBySpecificFilter(c, "rid", h.store.GetMealsByRestaurant)
}

// Mark an Item Special.
func (h *Handler) addSpecial(c echo.Context) (err error) {
	m := map[string]interface{}{}

	if err = c.Bind(&m); err != nil {
		return h.sendError(c, err)
	}

	if m["meal_id"] == nil || m["meal_id"] == "" {
		return h.sendError(c, err)
	}

	mid := m["meal_id"].(string)

	id, err := h.store.InsertSpecial(mid)

	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occurred     ",
		})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{
		"success": true,
		"id":      id,
	})
}

// getSpecials.
func (h *Handler) getSpecials(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetSpecialMeals)
}

// getSpicey
func (h *Handler) getSpicey(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetSpiceyMeals)
}

// getSpicey
func (h *Handler) searchMeal(c echo.Context) (err error) {
	keyword := c.QueryParam("keyword")
	return h.getFiltered(c, func() ([]interface{}, error) { return h.store.SearchMeal(keyword) })
}
