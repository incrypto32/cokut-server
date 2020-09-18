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

func FireAuthMiddleware(app *firebase.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.Background()

			// For testing purposes
			if true {
				return next(c)
			}

			if app == nil {
				return errors.New("Firebase App not initialized")
			}
			client, err := app.Auth(ctx)

			if err != nil {
				log.Println(err)
				return err
			}

			auth := c.Request().Header.Get("Authorization")
			idToken := strings.Replace(auth, "Bearer ", "", 1)
			token, err := client.VerifyIDToken(context.Background(), idToken)

			if err != nil {
				log.Println(err.Error())
				return c.JSON(http.StatusExpectationFailed, echo.Map{"success": false, "msg": "ID Token Expired"})
			}
			log.Println("Token Verified..!!")

			c.Set("token", token)
			return next(c)
		}
	}
}
