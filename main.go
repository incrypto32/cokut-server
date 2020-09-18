package main

import (
	"log"
	"os"

	"github.com/incrypt0/cokut-server/fire"
	"github.com/incrypt0/cokut-server/handler2"
	"github.com/incrypt0/cokut-server/handler2/middleware"
	"github.com/incrypt0/cokut-server/router"
	"github.com/incrypt0/cokut-server/store"
	"github.com/incrypt0/cokut-server/workers"
)

func main() {
	log.SetFlags(log.Llongfile)

	// Initialize Firebase
	app, err := fire.InitFire()

	if err != nil {
		log.Panic(err)
	}

	fireAuthMWare := middleware.FireAuthMiddleware(app)

	// Connect to mongo
	w := workers.New()
	s := store.NewStore("meals", "users", "orders", "restaurants", w)

	// echo instance
	r := router.New()

	// Main Echo Handler
	h := handler2.NewHandler(s, fireAuthMWare)
	h.Register(r)

	// Server Start
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("PORT is empty")
		PORT = "4000"
	}

	r.Logger.Fatal(r.Start(":" + PORT))
}
