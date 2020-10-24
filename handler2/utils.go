package handler2

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/incrypt0/cokut-server/brokers/myerrors"
	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add an item .
func (h *Handler) Add(c echo.Context, r models.Model, f func(r models.Model) (interface{}, error)) (err error) {
	log.Println(r.GetModelData())

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

// Add an item .
func (h *Handler) AddOrder(c echo.Context, r models.Model, f func(r models.Model) (interface{}, error)) (err error) {
	if err = c.Bind(r); err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	order, err := f(r)

	if err != nil {
		log.Println(err)

		return h.sendMessageWithFailure(c, "Order Not Validated", myerrors.ErrOrderNotValidatedCode)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"order":   order,
	})
}

func (h *Handler) getFiltered(c echo.Context, f ManyResultFunc) (err error) {
	l, err := f()

	if err != nil {
		return h.sendError(c, err)
	}

	if len(l) == 0 {
		return h.sendMessageWithFailure(c, "Nothing Here :(", myerrors.ErrNoRecordsCode)
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
		"code":    myerrors.ErrBasicCode,
		"msg":     "An error occurred",
	})
}

func (h *Handler) sendMessageWithFailure(c echo.Context, msg string, code int) error {
	return c.JSON(http.StatusExpectationFailed, echo.Map{
		"success": false,
		"msg":     msg,
		"code":    code,
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

	l, err := f(m[filter][0])

	if err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	if len(l) == 0 {
		return h.sendMessageWithFailure(c, "Nothing found there", myerrors.ErrNoRecordsCode)
	}

	return c.JSON(http.StatusOK, l)
}

func (h *Handler) parseRestaurantForm(c echo.Context) (pid primitive.ObjectID, r models.Restaurant, err error) {
	form, err := c.FormParams()
	if err != nil {
		log.Println(err)

		return pid, r, err
	}

	latitude, err := strconv.ParseFloat(form["latitude"][0], 64)

	if err != nil {
		log.Println(err)

		return pid, r, err
	}

	longitude, err := strconv.ParseFloat(form["longitude"][0], 64)

	if err != nil {
		log.Println(err)

		return pid, r, err
	}

	pid = primitive.NewObjectID()

	if err != nil {
		log.Println(err)

		return pid, r, err
	}

	location := models.Location{Latitude: latitude, Longitude: longitude}

	r = models.Restaurant{ID: pid,
		Name:     form["name"][0],
		Address:  form["address"][0],
		Type:     form["type"][0],
		Closed:   utils.NewBool(true),
		Location: &location}

	return pid, r, err
}
