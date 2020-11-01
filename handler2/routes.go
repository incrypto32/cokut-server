package handler2

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) registerAPI(api *echo.Group) {
	v1 := api.Group("/v1")
	h.registerAPIV1(v1)
}

func (h *Handler) registerAPIV1(api *echo.Group) {
	api.GET("/test", h.routeTestV1)

	// User
	user := api.Group("/user")
	user.GET("", h.getUser)
	user.GET("/orders", h.getUserOrders)
	user.POST("/register", h.registerUser)
	user.POST("/address", h.addAddress)
	user.DELETE("/address", h.removeAddress)

	// Restaurants
	restaurants := api.Group("/restaurants")
	restaurants.GET("", h.getAllRestaurants)
	restaurants.GET("/regular", h.getAllRegularRestaurants)
	restaurants.GET("/homemade", h.getHomeMadeRestaurants)

	restaurants.POST("/form", h.addRestaurantForm)
	restaurants.POST("/status", h.changeRestaurantStatus)
	restaurants.DELETE("/delete", h.deleteRestaurant)

	orders := api.Group("/order")
	orders.POST("", h.addOrder)
	orders.POST("/calculate", h.calculateOrder)

	// Meals
	meals := api.Group("/meals")
	meals.GET("", h.getMeals)
	meals.GET("/specials", h.getSpecials)
	meals.GET("/spicey", h.getSpicey)
	meals.GET("/search", h.searchMeal)
}

// The Admin Api .
func (h *Handler) registerAdmin(a *echo.Group) {
	a.GET("/test", h.routeTestAdmin)

	meal := a.Group("/meal")
	meal.POST("", h.addMeal)
	meal.DELETE("", h.deleteMeal)

	store := a.Group("/restaurant")
	store.POST("", h.addRestaurantForm)
	store.POST("/status", h.changeRestaurantStatus)
	store.DELETE("", h.deleteRestaurant)

	orders := a.Group("/orders")
	orders.GET("", h.getOrdersPaginated)
	orders.GET("/paginated", h.getOrders)

	a.POST("/special", h.addSpecial)
}

// The Admin Api .
func (h *Handler) registerUtils(u *echo.Group) {
	u.POST("/checkphone", h.checkUserPhoneExistence)
	u.POST("/checkuser", h.checkUserExistence)
	u.POST("/checkgid", h.checkUserExistence)
	u.POST("/getuser", h.getUser)
}

// Register this method registers a new group with handler .
func (h *Handler) Register(e *echo.Echo) {
	// Index Handler
	e.Use(middleware.CORS())
	e.GET("/", h.index)

	// Groups .
	api := e.Group("/api")

	api.GET("", h.routeTestAPI)

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

func (h *Handler) routeTestAPI(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"success": true, "msg": "HMM"})
}
