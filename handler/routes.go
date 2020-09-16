package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	a := v1.Group("/admin")
	a.POST("/addrest", h.addRestaurant)
	a.POST("/additem", h.addMeal)
	a.POST("/addspecial", h.addSpecial)
	a.POST("/order", h.addOrder)
	a.GET("/allrest", h.getAllRestaurants)
	a.GET("/getmeals", h.getMeals)
	a.GET("/getspecials", h.getSpecials)
	a.GET("/getspicey", h.getSpicey)
	a.GET("/gethome", h.getHomeMadeRestaurants)
	a.GET("/getorders", h.getOrders)
	a.GET("/getuserorders", h.getUserOrders)
}

func (h *Handler) Index(e *echo.Echo) {

	e.GET("/", h.index)

	// e.POST("/checkphone", h.checkUser)
}
