package handler2

import (
	"log"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) registerUser(c echo.Context) (err error) {
	uid := c.Get("uid")
	log.Println(uid)

	u, err := h.store.CheckUserExistenceByUID(uid.(string))

	if err != nil {
		log.Println(err)
		return h.sendError(c)
	}

	if u {
		return h.sendMessageWithFailure(c, "DETAILS_EXIST")
	}

	r := new(models.User)

	var id string
	if err := c.Bind(&r); err != nil {
		return h.sendError(c)
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
		log.Println(err)
		if err.Error() == "NIL" {
			err = nil

		} else {
			return h.sendError(c)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"user":    l,
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

func (h *Handler) checkUserExistenceByGID(c echo.Context) (err error) {

	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		log.Println(err)
		return h.sendError(c)
	}
	if m["gid"] == nil || m["gid"] == "" {
		return h.sendError(c)
	}

	exist, err := h.store.CheckUserPhoneExistenceByGID(m["gid"].(string))

	if err != nil && err.Error() != "NIL" {
		return h.sendError(c)
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
		return h.sendError(c)
	}
	if m["phone"] == nil || m["phone"] == "" {
		return h.sendError(c)
	}
	if m["email"] == nil || m["email"] == "" {
		return h.checkUserPhoneExistence(c)
	}

	exist, err := h.store.CheckUserExistence(m["phone"].(string), m["email"].(string))

	if err != nil && err.Error() != "NIL" {
		return h.sendError(c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"exist":   exist,
	})

}
