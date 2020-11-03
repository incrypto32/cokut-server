package handler2

import (
	"log"
	"net/http"
	"strconv"

	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/utils"
	"github.com/labstack/echo/v4"
)

// Create a new order .
func (h *Handler) addOrder(c echo.Context) (err error) {
	r := new(models.Order)
	r.UID = c.Get("uid").(string)

	return h.AddOrder(c, r, func(r models.Model) (interface{}, error) {
		return h.store.CreateOrder(r.(*models.Order), false)
	})
}

// Create a new order .
func (h *Handler) calculateOrder(c echo.Context) (err error) {
	o := new(models.Order)
	o.UID = c.Get("uid").(string)

	return h.AddOrder(c, o, func(r models.Model) (interface{}, error) {
		return h.store.CreateOrder(r.(*models.Order), true)
	})
}

func (h *Handler) getUserOrders(c echo.Context) (err error) {
	orders, err := h.store.GetOrdersByUser(c.Get("uid").(string))
	if err != nil {
		return h.sendError(c, err)
	}

	log.Println(utils.ModelToString(orders))

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"orders":  orders,
	})
}

func (h *Handler) getOrdersPaginated(c echo.Context) (err error) {
	limit, err := strconv.Atoi(c.QueryParam("limit"))

	if err != nil {
		limit = 500
	}

	page, err := strconv.Atoi(c.QueryParam("page"))

	if err != nil {
		page = 1
	}

	orders, err := h.store.GetPaginatedOrders(limit, page)

	if err != nil {
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"orders":  orders,
		"limit":   limit,
		"page":    page,
	})
}
