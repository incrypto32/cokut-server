package handler

import (
	// "fmt"
	"net/http"

	// "github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) routesHandler(e *echo.Echo) {
	// routes
	e.GET("/", h.index)

	// e.POST("/checkphone", h.checkUser)

}

// Index handler
func (h *Handler) index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "Hello,World")
}

// // Get Users Details
// func (h *Handler) GetUserDetails(c echo.Context) (err error) {
// 	u := new(models.User)

// 	if err = c.Bind(u); err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusExpectationFailed, echo.Map{
// 			"success": false,
// 			"msg":     err,
// 		})
// 	}

// 	u.ID = primitive.NewObjectID()

// 	id, err := models.InsertUser(u)

// 	if err != nil {
// 		return c.JSON(http.StatusExpectationFailed, echo.Map{
// 			"success": false,
// 			"msg":     err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusExpectationFailed, echo.Map{
// 		"success": true,
// 		"id":      id,
// 	})
// }

// // Check if a user exists with phone and email
// func (h *Handler) checkUser(c echo.Context) (err error) {
// 	fmt.Println("Authentication : ", c.Request().Header.Get("Authentication"))

// 	phone := c.FormValue("phone")

// 	if models.CheckUser(phone) {
// 		return c.JSON(http.StatusAccepted, echo.Map{
// 			"success": true,
// 			"exist":   true,
// 		})
// 	}
// 	return c.JSON(http.StatusAccepted, echo.Map{
// 		"success": true,
// 		"exist":   false,
// 	})
// }
