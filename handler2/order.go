package handler2

import (
	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Create a new order .
func (h *Handler) addOrder(c echo.Context) (err error) {
	r := new(models.Order)
	r.UID = c.Get("uid").(string)

	return h.AddOrder(c, r, func(r models.Model) (interface{}, error) {
		return h.store.CreateOrder(r.(*models.Order))
	})
}

func (h *Handler) getOrders(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllOrders)
}

func (h *Handler) getUserOrders(c echo.Context) (err error) {
	c.QueryParams().Add("uid", c.Get("uid").(string))

	return h.getBySpecificFilter(c, "uid", h.store.GetOrdersByUser)
}
