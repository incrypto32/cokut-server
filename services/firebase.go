package services

import (
	"context"

	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"google.golang.org/api/option"
)

var app *firebase.App

func InitFire() (*firebase.App, error) {
	var err error
	opt := option.WithCredentialsFile("./key.json")

	keys := os.Getenv("FIREBASE_KEYS")
	if keys != "" {
		fmt.Println("Reading from ENV")
		opt = option.WithCredentialsJSON([]byte(keys))
	}

	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, err
}

func FireAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.Background()

			if app == nil {
				return errors.New("Firebase App not initialized")
			}
			client, err := app.Auth(ctx)

			if err != nil {
				log.Error(err)
				return err
			}

			auth := c.Request().Header.Get("Authorization")
			idToken := strings.Replace(auth, "Bearer ", "", 1)
			token, err := client.VerifyIDToken(context.Background(), idToken)

			if err != nil {
				log.Error(err.Error())
				return c.JSON(http.StatusExpectationFailed, echo.Map{"success": false, "msg": "ID Token Expired"})
			}
			fmt.Println("Token Verified..!!")

			c.Set("token", token)
			return next(c)
		}
	}
}
