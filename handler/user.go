package handler

import (
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) registerUser(c echo.Context) (err error) {

	r := new(models.User)
	return h.Add(c, r, func(r models.Model) (string, error) {
		return h.userStore.Insert(r.(*models.User))
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

	exist := h.userStore.CheckUserExistence(m["phone"].(string))

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"exist":   exist,
	})

}
