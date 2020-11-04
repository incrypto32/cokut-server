package fire

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitFire() (*firebase.App, error) {
	var err error

	opt := option.WithCredentialsFile("./key.json")

	keys := os.Getenv("FIREBASE_KEYS")
	if keys != "" {
		log.Println("Reading from ENV")

		opt = option.WithCredentialsJSON([]byte(keys))
	}

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, err
}
