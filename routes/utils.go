package routes

import (
	"fmt"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/tester"
	"github.com/labstack/echo/v4"
)

// Add an item
func Add(c echo.Context, r models.Model, f func(r models.Model) (string, error)) (err error) {

	if err = c.Bind(r); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "An Error Occured",
		})
	}

	id, err := f(r)

	if err != nil {
		fmt.Println(err)
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

func getFiltered(c echo.Context, f func() ([]interface{}, error)) (err error) {
	tester.Tester()
	l, err := f()
	tester.Tester()
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
			"msg":     "Nothing found here :(",
		})
	}
	return c.JSON(http.StatusOK, l)

}
