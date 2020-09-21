package handler2

import (
	"log"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Create a new order
func (h *Handler) addOrder(c echo.Context) (err error) {
	r := new(models.Order)
	r.UID = "UID_HERE"

	return h.Add(c, r, func(r models.Model) (string, error) {
		return h.store.CreateOrder(r.(*models.Order))
	})
}

func (h *Handler) getOrders(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllOrders)
}

func (h *Handler) getUserOrders(c echo.Context) (err error) {
	m := map[string]interface{}{}
	if err = c.Bind(&m); err != nil {
		log.Println(err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occurred     ",
		})
	}

	if m["uid"] == nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occurred     ",
		})
	}

	l, err := h.store.GetOrdersByUser(m["uid"].(string))

	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occurred     ",
		})
	}

	if len(l) == 0 {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "User dont have any orders",
		})
	}

	return c.JSON(http.StatusOK, l)
}
