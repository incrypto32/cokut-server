package handler2

import (
	"log"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

// Add an item
func (h *Handler) Add(c echo.Context, r models.Model, f func(r models.Model) (string, error)) (err error) {
	if err = c.Bind(r); err != nil {
		log.Println(err)
		return h.sendError(c)
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

func (h *Handler) getFiltered(c echo.Context, f ManyResultFunc) (err error) {
	l, err := f()

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
			"msg":     "Nothing found here :(",
		})
	}

	return c.JSON(http.StatusOK, l)
}

func (h *Handler) sendError(c echo.Context) error {
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
	m := map[string]interface{}{}

	if err = c.Bind(&m); err != nil {
		log.Println(err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occurred     ",
		})
	}

	if m[filter] == nil {
		return h.sendError(c)
	}

	l, err := f(m[filter].(string))

	if err != nil {
		log.Println(err)

		return h.sendError(c)
	}

	if len(l) == 0 {
		return h.sendMessageWithFailure(c, "Nothing found there")
	}

	return c.JSON(http.StatusOK, l)
}
