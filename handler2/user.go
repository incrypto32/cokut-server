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
		return h.sendMessageWithFailure(c, err.Error())
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

	if _, err := h.store.AddUserAddress(c.Get("uid").(string), a); err != nil {
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}

func (h *Handler) checkUserPhoneExistence(c echo.Context) (err error) {
	m := echo.Map{}

	if err := c.Bind(&m); err != nil {
		return h.sendError(c, err)
	}

	if m["phone"] == nil || m["phone"] == "" {
		return h.sendError(c, err)
	}

	exist, err := h.store.CheckUserPhoneExistence(m["phone"].(string))
	if err != nil {
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"exist":   exist,
	})
}

// func (h *Handler) checkUserExistenceByGID(c echo.Context) (err error) {
// 	m := echo.Map{}

// 	if err := c.Bind(&m); err != nil {
// 		log.Println(err)
// 		return h.sendError(c,err)
// 	}

// 	if m["gid"] == nil || m["gid"] == "" {
// 		return h.sendError(c,err)
// 	}

// 	exist, err := h.store.CheckUserPhoneExistenceByGID(m["gid"].(string))

// 	if err != nil && err.Error() != "NIL" {
// 		return h.sendError(c,err)
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"success": true,
// 		"exist":   exist,
// 	})
// }

func (h *Handler) checkUserExistence(c echo.Context) (err error) {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		log.Println(err)
		return h.sendError(c, err)
	}

	if m["phone"] == nil || m["phone"] == "" {
		return h.sendError(c, err)
	}

	if m["email"] == nil || m["email"] == "" {
		return h.checkUserPhoneExistence(c)
	}

	exist, err := h.store.CheckUserExistence(m["phone"].(string), m["email"].(string))

	if err != nil && err.Error() != "NIL" {
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"exist":   exist,
	})
}
