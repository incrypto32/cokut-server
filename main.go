package main

import (
	"log"
	"os"

	"github.com/incrypt0/cokut-server/cokutbot"
	"github.com/incrypt0/cokut-server/fire"
	"github.com/incrypt0/cokut-server/handler2"
	"github.com/incrypt0/cokut-server/handler2/middleware"
	"github.com/incrypt0/cokut-server/myerrors"
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
	adminChecker := middleware.AdminCheckMiddleware()
	// New My Error
	e := myerrors.New()
	// Connect to mongo
	w := workers.New(e)

	// New Cokut Bot
	cbot, err := cokutbot.NewCokutBot()
	if err != nil {
		log.Panic("BOT FAILED")
	}

	s := store.NewStore("meals", "users", "orders", "restaurants", w, cbot, e)

	// echo instance
	r := router.New()

	// Main Echo Handler
	h := handler2.NewHandler(s, fireAuthMWare, adminChecker, e)
	h.Register(r)

	// Server Start
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("PORT is empty")

		PORT = "4000"
	}

	r.Logger.Fatal(r.Start(":" + PORT))
}
