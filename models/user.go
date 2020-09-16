package models

import (
	"context"
	"fmt"

	"github.com/incrypt0/cokut-server/services"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx context.Context = context.Background()

type User struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UID     string             `json:"uid,omitempty" bson:"uid,omitempty"`
	Name    string             `json:"name" bson:"name" `
	Phone   string             `json:"phone" bson:"phone" `
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	Address []string           `json:"address,omitempty" bson:"address,omitempty"`
	Admin   bool               `json:"admin,omitempty" bson:"admin,omitempty"`
}

// Prints Model Data in String
func (u *User) GetModelData() string {
	return services.PrintModel(u)
}

// Real Validation
func (u *User) Validate() error {
	fmt.Println("uid is : ", u.UID)
	if (u.Name == "") || (len(u.Phone) < 10) || u.UID == "" {
		return errors.New("Not Validated")
	}
	return nil
}

// Basic Validate
func (u *User) ValidateBasic() error {

	if (u.Name == "") || (len(u.Phone) < 10) {
		return errors.New("Enter Valid Details")
	}
	return nil
}

// // User Existence Middleware
// func UserExistenceMiddleWare() echo.MiddlewareFunc {
// 	handler := func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) (err error) {
// 			u := new(User)
// 			var token *auth.Token = c.Get("token").(*auth.Token)
// 			r := services.C.UserCollection.FindOne(context.Background(), bson.D{{Key: "uid", Value: token.UID}})

// 			if r.Err() != mongo.ErrNoDocuments {
// 				return c.JSON(http.StatusExpectationFailed, echo.Map{
// 					"success": false,
// 					"msg":     "User Not registered",
// 					"code":    "NOT_REGISTERED",
// 				})
// 			}

// 			if r.Err() != nil {
// 				return errors.New("Error")
// 			}
// 			if err = r.Decode(u); err != nil {
// 				log.Panic(err)
// 				return err
// 			}
// 			c.Set("user", u)
// 			return next(c)
// 		}

// 	}
// 	return handler
// }

// // Function to insert users into userCollection
// func InsertUser(u *User) (id string, err error) {
// 	//  Getting the user colection
// 	c := services.C.UserCollection

// 	// Basic Validation
// 	if err = u.ValidateBasic(); err != nil {
// 		return id, err
// 	}

// 	// Check if email is null
// 	if u.Email != "" {
// 		e := c.FindOne(ctx, bson.D{
// 			{Key: "email", Value: u.Email},
// 		})

// 		if e.Err() != mongo.ErrNoDocuments {
// 			err = errors.New("Email Address is already associated with another account")
// 			return id, err
// 		}
// 	}

// 	// Check if phone is null
// 	if u.Phone != "" {
// 		r := c.FindOne(ctx, bson.D{
// 			{Key: "phone", Value: u.Phone},
// 		})

// 		if r.Err() != mongo.ErrNoDocuments {
// 			err = errors.New("Phone number is already associated with another account")
// 			return id, err
// 		}
// 	}

// 	if err != nil {
// 		return id, err
// 	}

// 	return services.Add(c, u)
// }

// // Check User existence
// func CheckUserExistence(phone string) bool {
// 	var val bool = false
// 	fmt.Println("CheckUser called with phone : ", phone)
// 	c := services.C.UserCollection
// 	r := c.FindOne(ctx, bson.D{
// 		{Key: "phone", Value: phone},
// 	})

// 	if r.Err() != mongo.ErrNoDocuments {
// 		val = true
// 	}
// 	fmt.Println(phone, " exists : ", val)
// 	return val

// }

// // Admin check middleware
// func AdminCheck() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			if true {
// 				return next(c)
// 			}

// 			u := new(User)
// 			uid := c.Request().Header.Get("Uid")
// 			fmt.Println(uid)

// 			r := services.C.UserCollection.FindOne(ctx, bson.D{
// 				{Key: "uid", Value: uid},
// 			})

// 			if r.Err() == mongo.ErrNoDocuments {
// 				fmt.Println("no documents")
// 				return c.JSON(http.StatusUnauthorized, echo.Map{
// 					"success": false,
// 					"msg":     "User Unauthorized",
// 				})
// 			}

// 			if err := r.Decode(u); err != nil {
// 				return c.JSON(http.StatusUnauthorized, echo.Map{
// 					"success": false,
// 					"msg":     "User Unauthorized",
// 				})
// 			}
// 			if u.Admin {
// 				return next(c)
// 			} else {
// 				return c.JSON(http.StatusUnauthorized, echo.Map{
// 					"success": false,
// 					"msg":     "User Unauthorized",
// 				})
// 			}

// 		}
// 	}

// }
