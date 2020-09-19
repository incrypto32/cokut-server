package handler2

import (
	"log"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) registerUser(c echo.Context) (err error) {

	r := new(models.User)
	return h.Add(c, r, func(r models.Model) (string, error) {
		return h.store.InsertUser(r.(*models.User))
	})
}

func (h *Handler) checkUserPhoneExistence(c echo.Context) (err error) {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	if m["phone"] == nil || m["phone"] == "" {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	exist, err := h.store.CheckUserPhoneExistence(m["phone"].(string))
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"exist":   exist,
	})

}

func (h *Handler) checkUserExistence(c echo.Context) (err error) {

	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		log.Println(err)
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	if m["phone"] == nil || m["phone"] == "" {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	if m["email"] == nil || m["email"] == "" {
		return h.checkUserPhoneExistence(c)
	}

	exist, err := h.store.CheckUserExistence(m["phone"].(string), m["email"].(string))

	if err != nil && err.Error() != "NIL" {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"exist":   exist,
	})

}
