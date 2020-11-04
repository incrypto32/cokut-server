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

// Change Order Status .
func (h *Handler) changeOrderStatus(c echo.Context) (err error) {
	params := make(map[string]interface{})

	if err = c.Bind(&params); err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	o, err := h.store.ChangeOrderStatus(params["id"].(string), int(params["statusCode"].(float64)))
	if err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "order": o})
}

func (h *Handler) getUserOrders(c echo.Context) (err error) {
	orders, err := h.store.GetOrdersByUser(c.Get("uid").(string))
	if err != nil {
		return h.sendError(c, err)
	}

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
