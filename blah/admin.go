package blah

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/incrypt0/cokut-server/models"

// 	"github.com/labstack/echo/v4"
// )

func Blah() {}

// func Admin(g *echo.Group) {
// 	g.Use(models.AdminCheck())
// 	g.POST("/addrest", addRestaurant)
// 	g.POST("/additem", addMeal)
// 	g.POST("/addspecial", addSpecial)
// 	g.POST("/order", addOrder)
// 	g.GET("/allrest", getAllRestaurants)
// 	g.GET("/getmeals", getMeals)
// 	g.GET("/getspecials", getSpecials)
// 	g.GET("/getspicey", getSpicey)
// 	g.GET("/gethome", getHome)
// 	g.GET("/getorders", getOrders)
// 	g.GET("/getuserorders", getUserOrders)
// }

// // Add a meal to the db
// func addMeal(c echo.Context) (err error) {
// 	r := new(models.Meal)
// 	return Add(c, r, func(r models.Model) (string, error) {
// 		return models.InsertMeal(r.(*models.Meal))
// 	})
// }

// // Add a single restaurant
// func addRestaurant(c echo.Context) (err error) {
// 	r := new(models.Restaurant)
// 	return Add(c, r, func(r models.Model) (string, error) {
// 		return models.InsertRestaurant(r.(*models.Restaurant))
// 	})
// }

// // Create a new order
// func addOrder(c echo.Context) (err error) {
// 	r := new(models.Order)
// 	return Add(c, r, func(r models.Model) (string, error) {
// 		return models.InsertOrder(r.(*models.Order), "blah4")
// 	})
// }

// // Mark an Item Special
// func addSpecial(c echo.Context) (err error) {
// 	m := map[string]interface{}{}

// 	if err = c.Bind(&m); err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}
// 	if m["meal_id"] == nil || m["meal_id"] == "" {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}
// 	mid := m["meal_id"].(string)

// 	id, err := models.InsertSpecial(mid)

// 	if err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}

// 	return c.JSON(http.StatusInternalServerError, echo.Map{
// 		"success": true,
// 		"id":      id.Hex(),
// 	})

// }

// // Get all restaurants in the db
// func getAllRestaurants(c echo.Context) (err error) {

// 	l, err := models.GetAllRestaurants()

// 	if err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, l)
// }

// // Get all meals from the database with the given resaurant ID
// func getMeals(c echo.Context) (err error) {

// 	m := map[string]interface{}{}

// 	if err = c.Bind(&m); err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}

// 	if m["rid"] == nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}

// 	l, err := models.GetMeals(m["rid"].(string))

// 	fmt.Println(l)
// 	if err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}

// 	if len(l) <= 0 {
// 		return c.JSON(http.StatusExpectationFailed, echo.Map{
// 			"success": false,
// 			"msg":     "Restaurant dont exist",
// 		})
// 	}
// 	return c.JSON(http.StatusOK, l)

// }

// // getSpecials
// func getSpecials(c echo.Context) (err error) {
// 	return getFiltered(c, models.GetSpecials)
// }

// // getSpicey
// func getSpicey(c echo.Context) (err error) {
// 	return getFiltered(c, models.GetSpicey)
// }

// //get Home
// func getHome(c echo.Context) (err error) {
// 	return getFiltered(c, models.GetHomeMade)
// }

// // // Get all orders for admin
// // func getAllOrders(c echo.Context) (err error) {

// // 	l, err := models.GetOrders()

// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return c.JSON(http.StatusInternalServerError, echo.Map{
// // 			"success": false,
// // 			"msg":     "An error occured",
// // 		})
// // 	}

// // 	return c.JSON(http.StatusOK, l)
// // }

// //get Home
// func getOrders(c echo.Context) (err error) {

// 	return getFiltered(c, models.GetOrders)
// }

// // Get all orders for admin
// func getUserOrders(c echo.Context) (err error) {

// 	m := map[string]interface{}{}

// 	if err = c.Bind(&m); err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}

// 	if m["uid"] == nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}

// 	l, err := models.GetUserOrders(m["uid"].(string))

// 	fmt.Println(l)

// 	if err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"success": false,
// 			"msg":     "An error occured",
// 		})
// 	}

// 	if len(l) <= 0 {
// 		return c.JSON(http.StatusExpectationFailed, echo.Map{
// 			"success": false,
// 			"msg":     "User dont have any orders",
// 		})
// 	}
// 	return c.JSON(http.StatusOK, l)
// }
