package handler2

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/incrypt0/cokut-server/brokers/myerrors"
	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add a single restaurant
func (h *Handler) addRestaurant(c echo.Context) (err error) {
	r := new(models.Restaurant)

	return h.Add(c, r, func(r models.Model) (interface{}, error) {
		return h.store.InsertRestaurant(r.(*models.Restaurant))
	})
}

func (h *Handler) addRestaurantForm(c echo.Context) (err error) {
	form, err := c.FormParams()
	latitude, err := strconv.ParseFloat(form["latitude"][0], 64)
	longitude, err := strconv.ParseFloat(form["longitude"][0], 64)

	pid := primitive.NewObjectID()

	if err != nil {
		log.Println(err)
		return h.sendError(c, err)
	}

	location := models.Location{Latitude: latitude, Longitude: longitude}

	r := models.Restaurant{ID: pid, Name: form["name"][0], Address: form["address"][0], Closed: utils.NewBool(true), Location: &location}

	log.Println(r.GetModelData())

	if err != nil {
		log.Println(err)
		return h.sendError(c, err)
	}

	file, err := c.FormFile("file")

	if err != nil {
		log.Println(err)
		return h.sendMessageWithFailure(c, "Please upload a vallid file", myerrors.FileUploadErrorCode)
	}

	log.Println(file.Filename)

	src, err := file.Open()
	if err != nil {
		log.Println(err)
		return h.sendMessageWithFailure(c, "Please upload a vallid file", myerrors.FileUploadErrorCode)
	}
	defer src.Close()

	// Destination

	dst, err := os.Create("./files/restaurants/" + pid.Hex() + ".png")
	if err != nil {
		log.Println(err)
		return h.sendMessageWithFailure(c, "Please upload a vallid file", myerrors.FileUploadErrorCode)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Println(err)
		return h.sendMessageWithFailure(c, "Please upload a vallid file", myerrors.FileUploadErrorCode)
	}

	return c.JSON(http.StatusOK, form)
}

func (h *Handler) changeRestaurantStatus(c echo.Context) (err error) {
	params := make(map[string]interface{})

	if err = c.Bind(&params); err != nil {
		log.Println(err)
		return h.sendError(c, err)
	}

	id := params["id"].(string)
	value := params["closed"].(bool)

	r := models.Restaurant{Closed: utils.NewBool(value)}

	result, err := h.store.UpdateRestaurantStatus(id, r)

	if err != nil {
		log.Println(err)
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, result)
}
func (h *Handler) deleteRestaurant(c echo.Context) (err error) {

	log.Println(c.QueryParam("id"))

	a, err := h.store.DeleteRestaurant(c.QueryParam("id"))

	if err != nil {
		log.Println(err)
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, a)
}

// Get all restaurants in the db
func (h *Handler) getAllRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllRestaurants)
}

// Get all restaurants in the db
func (h *Handler) getAllRegularRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllRegularRestaurants)
}

//get Home
func (h *Handler) getHomeMadeRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllHomeMade)
}
