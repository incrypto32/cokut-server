package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
)

// FireAuthMiddleware is a middlerware which handles the authorisation checking for users
func FireAuthMiddleware(app *firebase.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.Background()
			client, err := app.Auth(ctx)
			if err != nil {
				log.Println(err)
				return err
			}

			// if true {

			// 	if false {
			// 		c.Set("uid", "u1")
			// 		return next(c)
			// 	}
			// }

			if app == nil {
				return errors.New("Firebase App not initialized")
			}

			auth := c.Request().Header.Get("Authorization")

			idToken := strings.Replace(auth, "Bearer ", "", 1)

			token, err := client.VerifyIDToken(context.Background(), idToken)

			if err != nil {
				log.Println(err.Error())
				return c.JSON(http.StatusExpectationFailed, echo.Map{"success": false, "msg": "ID Token Expired"})
			}
			log.Println("Token Verified..!!")

			c.Set("uid", token.UID)
			return next(c)
		}
	}
}
