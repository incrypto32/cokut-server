package main

import (
	"fmt"
	"log"
	"os"

	"github.com/incrypt0/cokut-server/handler"
	"github.com/incrypt0/cokut-server/router"
	"github.com/incrypt0/cokut-server/services"
)

func main() {
	// Initialize Firebase
	if _, err := services.InitFire(); err != nil {
		log.Panic(err)
	}

	// Connect to mongo
	services.ConnectMongo()

	// echo instance
	r := router.New()
	v1 := r.Group("/api")
	h := handler.NewHandler()
	h.Register(v1)

	// Server Start
	PORT := os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("PORT is empty")
		PORT = "4000"
	} else {
		PORT = "4000"
	}

	r.Logger.Fatal(r.Start(":" + PORT))
}
