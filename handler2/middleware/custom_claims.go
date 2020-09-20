package middleware

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

// CustomClaimRegister middleware
func CustomClaimRegister(app *firebase.App) func(string, map[string]interface{}) error {

	return func(uid string, claims map[string]interface{}) error {

		ctx := context.Background()
		client, err := app.Auth(ctx)
		if err != nil {
			log.Printf("error getting Auth client: %v\n", err)
			return err
		}

		err = client.SetCustomUserClaims(ctx, uid, claims)
		if err != nil {
			log.Printf("error setting custom claims %v\n", err)
			return err
		}
		return nil

	}
}
