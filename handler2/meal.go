package handler2

import (
	"log"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Add a meal to the db
func (h *Handler) addMeal(c echo.Context) (err error) {
	r := new(models.Meal)
	return h.Add(c, r, func(r models.Model) (string, error) {
		return h.store.InsertMeal(r.(*models.Meal))
	})
}

// Get all meals from the database with the given resaurant ID
func (h *Handler) getMeals(c echo.Context) (err error) {

	m := map[string]interface{}{}

	if err = c.Bind(&m); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	if m["rid"] == nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	l, err := h.store.GetMealsByRestaurant(m["rid"].(string))

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	if len(l) <= 0 {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "Nothing found there",
		})
	}

	return c.JSON(http.StatusOK, l)

}

// Mark an Item Special
func (h *Handler) addSpecial(c echo.Context) (err error) {
	m := map[string]interface{}{}

	if err = c.Bind(&m); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	if m["meal_id"] == nil || m["meal_id"] == "" {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	mid := m["meal_id"].(string)

	id, err := h.store.InsertSpecial(mid)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{
		"success": true,
		"id":      id,
	})

}

// getSpecials
func (h *Handler) getSpecials(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetSpecialMeals)
}

// getSpicey
func (h *Handler) getSpicey(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetSpiceyMeals)
}
