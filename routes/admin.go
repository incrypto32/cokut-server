package routes

import (
	"fmt"
	"net/http"

	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Admin(g *echo.Group) {
	g.Use(models.AdminCheck())
	g.POST("/addrest", addRestaurant)
	g.POST("/additem", addMeal)
	g.POST("/addspecial", addSpecial)
	g.GET("/allrest", getAllRestaurants)
}

// Add a meal to the db
func addMeal(c echo.Context) (err error) {
	r := new(models.Meal)
	return Add(c, r, func(r models.Model) (primitive.ObjectID, error) {
		return models.InsertMeal(r.(*models.Meal))
	})
}

// Add a single restaurant
func addRestaurant(c echo.Context) (err error) {
	r := new(models.Restaurant)
	return Add(c, r, func(r models.Model) (primitive.ObjectID, error) {
		return models.InsertRestaurant(r.(*models.Restaurant))
	})
}

// Mark an Item Special
func addSpecial(c echo.Context) (err error) {
	m := map[string]interface{}{}

	if err = c.Bind(&m); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	if m["meal_id"] == nil || m["meal_id"] == "" {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}
	mid := m["meal_id"].(string)

	id, err := models.InsertSpecial(mid)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{
		"success": true,
		"id":      id.Hex(),
	})

}

// Get all restaurants in the db
func getAllRestaurants(c echo.Context) (err error) {

	l, err := models.GetAllRestaurants()
	fmt.Println()
	fmt.Println(l)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"msg":     "An error occured",
		})
	}

	return c.JSON(http.StatusOK, l)
}
