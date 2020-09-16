package handler

import (
	"fmt"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Add a meal to the db
func (h *Handler) addMeal(c echo.Context) (err error) {
	r := new(models.Meal)
	return h.Add(c, r, func(r models.Model) (string, error) {
		return models.InsertMeal(r.(*models.Meal))
	})
}

// Get all meals from the database with the given resaurant ID
func (h *Handler) getMeals(c echo.Context) (err error) {

	m := map[string]interface{}{}

	if err = c.Bind(&m); err != nil {
		fmt.Println(err)
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

	l, err := models.GetMeals(m["rid"].(string))

	fmt.Println(l)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	if len(l) <= 0 {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "Restaurant dont exist",
		})
	}
	return c.JSON(http.StatusOK, l)

}

// Mark an Item Special
func (h *Handler) addSpecial(c echo.Context) (err error) {
	m := map[string]interface{}{}

	if err = c.Bind(&m); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	if m["meal_id"] == nil || m["meal_id"] == "" {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	mid := m["meal_id"].(string)

	id, err := models.InsertSpecial(mid)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{
		"success": true,
		"id":      id.Hex(),
	})

}

// getSpecials
func (h *Handler) getSpecials(c echo.Context) (err error) {
	return h.getFiltered(c, models.GetSpecials)
}

// getSpicey
func (h *Handler) getSpicey(c echo.Context) (err error) {
	return h.getFiltered(c, models.GetSpicey)
}
