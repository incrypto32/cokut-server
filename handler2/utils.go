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
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "An Error Occured",
		})
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

func (h *Handler) getFiltered(c echo.Context, f func() ([]interface{}, error)) (err error) {

	l, err := f()

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
			"msg":     "Nothing found here :(",
		})
	}
	return c.JSON(http.StatusOK, l)

}
