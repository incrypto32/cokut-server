package middleware

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

// CustomClaimRegister middleware.
func CustomClaimRegister(app *firebase.App) func(string, map[string]interface{}) error {
	log.Println("Inside Custom Claims Handler")

	return func(uid string, claims map[string]interface{}) error {
		log.Println("____________________Test 1")

		ctx := context.Background()
		client, err := app.Auth(ctx)

		if err != nil {
			log.Printf("error getting Auth client: %v\n", err)
			return err
		}

		log.Println("____________________Test 2")
		log.Println(claims)

		err = client.SetCustomUserClaims(ctx, uid, claims)

		log.Println("____________________Test 3")

		if err != nil {
			log.Printf("error setting custom claims %v\n", err)
			return err
		}

		log.Println("____________________Test 4")

		return nil
	}
}
