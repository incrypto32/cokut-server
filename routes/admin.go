package routes

import (
	"github.com/incrypt0/cokut-server/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Admin(g *echo.Group) {
	g.Use(models.AdminCheck())
	g.POST("/addrest", addRest)
	g.POST("/additem", addMeal)
}

// // Add restaurants
// func addRestaurant(c echo.Context) (err error) {
// 	r := new(models.Restuarant)

// 	if err = c.Bind(r); err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusExpectationFailed, echo.Map{
// 			"success": false,
// 			"msg":     "An Error Occured",
// 		})
// 	}

// 	id, err := models.InsertRestaurant(r)

// 	if err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusExpectationFailed, echo.Map{
// 			"success": false,
// 			"msg":     err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"success": true,
// 		"msg":     "pwoliyeee",
// 		"id":      id,
// 	})
// }

func addMeal(c echo.Context) (err error) {
	r := new(models.Meal)
	return Add(c, r, func(r models.Model) (primitive.ObjectID, error) {
		return models.InsertMeal(r.(*models.Meal))
	})
}

func addRest(c echo.Context) (err error) {
	r := new(models.Restuarant)
	return Add(c, r, func(r models.Model) (primitive.ObjectID, error) {
		return models.InsertRestaurant(r.(*models.Restuarant))
	})
}
