package services

import (
	"context"
	"encoding/json"
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
	return auth
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

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
		token, err := client.VerifyIDToken(ctx, idToken)

		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusExpectationFailed, echo.Map{"success": false, "msg": "ID Token Expired"})
		}
		fmt.Println("Token Verified..!!")

		c.Set("token", token)
		return next(c)
	}
}

func Test() {
	fmt.Println("Testing Fiire")
	fmt.Println(app)
	client, err := app.Auth(ctx)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Yogillya Amminye 1")
	}

	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjQ5YWQ5YmM1ZThlNDQ3OTNhMjEwOWI1NmUzNjFhMjNiNDE4ODA4NzUiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20va29zcy00ODhhYyIsImF1ZCI6Imtvc3MtNDg4YWMiLCJhdXRoX3RpbWUiOjE1OTk4MDE1MjQsInVzZXJfaWQiOiIzdkFRQkdLVFUyUU5oc2hhcTk0VWRHeHprUm8yIiwic3ViIjoiM3ZBUUJHS1RVMlFOaHNoYXE5NFVkR3h6a1JvMiIsImlhdCI6MTU5OTgwMTUyNCwiZXhwIjoxNTk5ODA1MTI0LCJwaG9uZV9udW1iZXIiOiIrOTE3MDM0MzIwNDQwIiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJwaG9uZSI6WyIrOTE3MDM0MzIwNDQwIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGhvbmUifX0.zHZ13FJeCYFF315CIcRen_JTf_a4yGgz-0Q-y_Ii9SkmvqYmdISctZdEv42BM2nPGNL9LGPZ0yRKCXe5lZA6b0sFgCaVHrIKhDClqYCze-lmMYX7z9LiVvjJHhi1Li0fYkgn3EeL5uzn5EeA_R2Vi3xR1EbgwWN8iZ08feWOrZ3mjOrfsHGtVPPBpbFUyZQkI82D8Rf59L99L3e9Sp6ZJHXCRCsdJJ_lQkEL57XNKzsB8a9Ujjw1BYqHQGZ5iyubqfdOe868DT5HSNFpOwUB9TCvq2f2bEisspIXpFW5gjljU1z2q76t1EroS-fyJqQ_YABEzbfCWQCvqAlhy7ddKw"
	token, err := client.VerifyIDToken(ctx, idToken)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Yogillya amminye")
	}
	b, _ := json.MarshalIndent(token, "", "  ")
	fmt.Println(string(b))
}
