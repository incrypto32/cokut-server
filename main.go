package main

import (
	"fmt"
	"log"
	"os"

	"github.com/incrypt0/cokut-server/fire"
	"github.com/incrypt0/cokut-server/handler2"
	"github.com/incrypt0/cokut-server/router"
	"github.com/incrypt0/cokut-server/store"
	"github.com/incrypt0/cokut-server/workers"
)

func main() {

	// Initialize Firebase
	_, err := fire.InitFire()

	if err != nil {
		log.Panic(err)
	}

	// Connect to mongo
	w := workers.New()
	s := store.NewStore("meals", "users", "orders", "restaurants", w)

	// echo instance
	r := router.New()

	h := handler2.NewHandler(s)
	h.Register(r)

	// Server Start
	PORT := os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("PORT is empty")
		PORT = "4000"
	}

	r.Logger.Fatal(r.Start(":" + PORT))
}
