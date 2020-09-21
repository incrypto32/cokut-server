package handler2

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) registerAPI(api *echo.Group) {
	v1 := api.Group("/v1")
	h.registerAPIV1(v1)
}

func (h *Handler) registerAPIV1(api *echo.Group) {
	log.Println("________API V1 Handler Initiated________")
	api.GET("/test", h.routeTestV1)
	api.GET("/getuser", h.getUser)
	api.GET("/getoutlets", h.getAllRestaurants)
	api.GET("/getmeals", h.getMeals)
	api.GET("/getspecials", h.getSpecials)
	api.GET("/getspicey", h.getSpicey)
	api.GET("/gethome", h.getHomeMadeRestaurants)
	api.GET("/getuserorders", h.getUserOrders)

	api.POST("/register", h.registerUser)
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

// The Admin Api
func (h *Handler) registerUtils(u *echo.Group) {
	u.POST("/checkphone", h.checkUserPhoneExistence)
	u.POST("/checkuser", h.checkUserExistence)
	u.POST("/checkgid", h.checkUserExistence)
	u.POST("/getuser", h.getUser)
}

// Register this method registers a new group with handler
func (h *Handler) Register(e *echo.Echo) {
	// Index Handler
	e.GET("/", h.index)

	// Groups
	api := e.Group("/api")
	admin := e.Group("/admin")
	u := e.Group("/utils")

	// middlewares
	api.Use(h.fireAuthMWare)

	// Register the routes
	h.registerUtils(u)
	h.registerAPI(api)
	h.registerAdmin(admin)
}

func (h *Handler) routeTestAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"success": true, "msg": "Admin Working Fine"})
}

func (h *Handler) routeTestV1(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"success": true, "msg": "V1 working fine"})
}
