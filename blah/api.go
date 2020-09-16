package blah

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/incrypt0/cokut-server/models"
// 	"github.com/incrypt0/cokut-server/services"
// 	"github.com/labstack/echo/v4"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

func Hi() {

}

// func Api(g *echo.Group) {
// 	g.Use(services.FireAuthMiddleware())
// 	g.POST("/register", RegisterUser)
// 	g.POST("/check", Check)

// }

// // Register a User
// func RegisterUser(c echo.Context) (err error) {
// 	u := new(models.User)

// 	if err = c.Bind(u); err != nil {
// 		fmt.Println(err)
// 		return c.JSON(http.StatusExpectationFailed, echo.Map{
// 			"success": false,
// 			"msg":     "An Error Occured",
// 		})
// 	}

// 	u.ID = primitive.NewObjectID()

// 	id, err := models.InsertUser(u)

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

// // Testing
// func Check(c echo.Context) (err error) {
// 	id := c.QueryParams().Get("id")
// 	u := new(models.User)
// 	p, _ := primitive.ObjectIDFromHex(id)
// 	r := services.C.UserCollection.FindOne(context.Background(), bson.D{
// 		{Key: "_id", Value: p},
// 	})
// 	if err = r.Decode(u); err != nil {
// 		fmt.Println("Errrrorr")
// 		fmt.Println(err)
// 		return c.JSON(http.StatusExpectationFailed, u)
// 	}
// 	return c.JSON(http.StatusOK, u)

// }
