package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) registerApi(api *echo.Group) {
	v1 := api.Group("/v1")
	h.registerApiV1(v1)
}

func (h *Handler) registerApiV1(api *echo.Group) {

	fmt.Println("________API Handler Initiated________")
	api.GET("/test", h.routeTestV1)
	api.GET("/getoutlets", h.getAllRestaurants)
	api.GET("/getmeals", h.getMeals)
	api.GET("/getspecials", h.getSpecials)
	api.GET("/getspicey", h.getSpicey)
	api.GET("/gethome", h.getHomeMadeRestaurants)
	api.GET("/getuserorders", h.getUserOrders)

	api.POST("/register", h.registerUser)
	api.POST("/checkphone", h.checkUserPhoneExistence)
	api.POST("/order", h.addOrder)

	a := api.Group("/admin")
	h.registerAdmin(a)

}

// The Admin Api
func (h *Handler) registerAdmin(a *echo.Group) {

	a.GET("/test", h.routeTestAdmin)
	a.GET("/getorders", h.getOrders)
	a.POST("/addrest", h.addRestaurant)
	a.POST("/additem", h.addMeal)
	a.POST("/addspecial", h.addSpecial)

}

func (h *Handler) Register(e *echo.Echo) {

	// Index Handler
	e.GET("/", h.index)

	// Groups
	api := e.Group("/api")
	admin := e.Group("/admin")

	// Register the routes
	h.registerApi(api)
	h.registerAdmin(admin)
}

func (h *Handler) routeTestAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"success": true, "msg": "Admin Working Fine"})
}

func (h *Handler) routeTestV1(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"success": true, "msg": "V1 working fine"})
}
