package handler2

import (
	"log"
	"net/http"
	"strconv"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Create a new order .
func (h *Handler) addOrder(c echo.Context) (err error) {
	r := new(models.Order)
	r.UID = c.Get("uid").(string)

	log.Println("1")

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

func (h *Handler) getOrdersPaginated(c echo.Context) (err error) {
	limit, err := strconv.Atoi(c.QueryParam("limit"))

	if err != nil {
		limit = 5
	}

	page, err := strconv.Atoi(c.QueryParam("page"))

	if err != nil {
		page = 1
	}

	log.Println(limit, page)

	orders, err := h.store.GetPaginatedOrders(limit, page)

	log.Println(len(orders))

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
