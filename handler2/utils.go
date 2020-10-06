package handler2

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Add an item
func (h *Handler) Add(c echo.Context, r models.Model, f func(r models.Model) (interface{}, error)) (err error) {
	if err = c.Bind(r); err != nil {
		log.Println(err)
		return h.sendError(c, err)
	}

	id, err := f(r)

	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"msg":     "pwoliyeee",
		"id":      id,
	})
}

// Add an item
func (h *Handler) AddOrder(c echo.Context, r models.Model, f func(r models.Model) (interface{}, error)) (err error) {
	if err = c.Bind(r); err != nil {
		log.Println(err)
		return h.sendError(c, err)
	}

	order, err := f(r)

	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"msg":     "pwoliyeee",
		"order":   order,
	})
}

func (h *Handler) getFiltered(c echo.Context, f ManyResultFunc) (err error) {
	l, err := f()

	if err != nil {
		return h.sendError(c, err)
	}

	if len(l) == 0 {
		return h.sendMessageWithFailure(c, "Nothing Here :(")
	}

	return c.JSON(http.StatusOK, l)
}

func (h *Handler) sendError(c echo.Context, err error) error {
	function, file, line, _ := runtime.Caller(1)
	trace := fmt.Sprintf("\nFile: %s  Function: %s Line: %d", file, runtime.FuncForPC(function).Name(), line)
	log.Println("Error Sent : ", err, trace)

	return c.JSON(http.StatusExpectationFailed, echo.Map{
		"success": false,
		"error":   true,
		"msg":     "An error occurred     ",
	})
}

func (h *Handler) sendMessageWithFailure(c echo.Context, msg string) error {
	return c.JSON(http.StatusExpectationFailed, echo.Map{
		"success": false,
		"msg":     msg,
	})
}

// Get all meals from the database with the given restaurant ID.
func (h *Handler) getBySpecificFilter(
	c echo.Context, filter string,
	f FilteredManyResultFunc) (err error) {
	m := c.QueryParams()

	if m[filter] == nil {
		return h.sendError(c, err)
	}

	log.Println(m[filter][0])

	l, err := f(m[filter][0])

	if err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	if len(l) == 0 {
		return h.sendMessageWithFailure(c, "Nothing found there")
	}

	return c.JSON(http.StatusOK, l)
}
