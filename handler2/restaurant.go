package handler2

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/incrypt0/cokut-server/brokers/myerrors"
	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// // Add a single restaurant .
// func (h *Handler) addRestaurant(c echo.Context) (err error) {
// 	r := new(models.Restaurant)

// 	return h.Add(c, r, func(r models.Model) (interface{}, error) {
// 		return h.store.InsertRestaurant(r.(*models.Restaurant))
// 	})
// }

func (h *Handler) addRestaurantForm(c echo.Context) (err error) {
	pid, r, err := h.parseRestaurantForm(c)

	if err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	if err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return h.sendMessageWithFailure(c, "Please upload a vallid file", myerrors.FileUploadErrorCode)
	}

	if err = h.handleFile(file, pid); err != nil {
		return h.sendMessageWithFailure(c, "Please upload a vallid file", myerrors.FileUploadErrorCode)
	}

	var result string

	if result, err = h.store.InsertRestaurant(&r); err != nil {
		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"id":      result,
	})
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

	result, err := h.store.UpdateRestaurant(id, r)

	if err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success":    true,
		"restaurant": result,
	})
}
func (h *Handler) handleFile(file *multipart.FileHeader, pid primitive.ObjectID) (err error) {
	if err != nil {
		log.Println(err)

		return err
	}

	src, err := file.Open()
	if err != nil {
		log.Println(err)

		return err
	}
	defer src.Close()

	// Destination

	dst, err := os.Create("./files/restaurants/" + pid.Hex() + ".png")
	if err != nil {
		log.Println(err)

		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Println(err)

		return err
	}

	if err != nil {
		log.Println(err)

		return err
	}

	return err
}

func (h *Handler) deleteRestaurant(c echo.Context) (err error) {
	a, err := h.store.DeleteRestaurant(c.QueryParam("id"))

	if err != nil {
		log.Println(err)

		return h.sendError(c, err)
	}

	if err := os.Remove("files/restaurants/" + c.QueryParam("id") + ".png"); err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success":  true,
		"ndeleted": a,
	})
}

// Get all restaurants in the db .
func (h *Handler) getAllRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllRestaurants)
}

// Get all restaurants in the db .
func (h *Handler) getAllRegularRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllRegularRestaurants)
}

// get Home .
func (h *Handler) getHomeMadeRestaurants(c echo.Context) (err error) {
	return h.getFiltered(c, h.store.GetAllHomeMade)
}
