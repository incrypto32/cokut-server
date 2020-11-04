package handler2

import (
	"log"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) registerUser(c echo.Context) (err error) {
	uid := c.Get("uid")

	r := new(models.User)
	r.UID = uid.(string)

	var id string

	if err := c.Bind(&r); err != nil {
		return h.sendError(c, err)
	}

	r.UID = uid.(string)
	if id, err = h.store.InsertUser(r); err != nil {
		log.Println(err)

		return h.sendMessageWithFailure(c, err.Error(), h.myerrors.ErrBasicCode)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"id":      id,
	})
}

func (h *Handler) getUser(c echo.Context) (err error) {
	uid := c.Get("uid")

	l, err := h.store.GetUser(uid.(string))

	if err != nil {
		if err.Error() == "NIL" {
			return c.JSON(http.StatusOK, echo.Map{
				"success": true,
				"exist":   false,
				"user":    nil,
			})
		}

		log.Println(err)

		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"exist":   true,
		"user":    l,
	})
}

func (h *Handler) addAddress(c echo.Context) (err error) {
	a := models.Address{}

	if err := c.Bind(&a); err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	user, err := h.store.AddUserAddress(c.Get("uid").(string), a)

	if err != nil {
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "user": user})
}

func (h *Handler) removeAddress(c echo.Context) (err error) {
	a := models.Address{}

	a.Title = c.QueryParam("title")

	user, err := h.store.RemoveUserAddress(c.Get("uid").(string), a)
	if err != nil {
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "user": user})
}
