package main

import (
	"fmt"
	"log"
	"os"

	"github.com/incrypt0/cokut-server/fire"
	"github.com/incrypt0/cokut-server/handler"
	"github.com/incrypt0/cokut-server/router"
	"github.com/incrypt0/cokut-server/stores"
	"github.com/incrypt0/cokut-server/workers"
)

func main() {

	// Initialize Firebase
	_, err := fire.InitFire()

	if err != nil {
		log.Panic(err)
	}

	// Connect to mongo
	db := workers.ConnectMongo()

	// Initialize refernce to required Collections
	users := db.Collection("users")
	meals := db.Collection("meals")
	restaurants := db.Collection("restaurants")
	orders := db.Collection("orders")

	userStore := stores.NewUserStore(users)
	mealsStore := stores.NewMealStore(meals, restaurants)
	restaurantStore := stores.NewRestaurantStore(restaurants)
	orderStore := stores.NewOrderStore(orders, restaurants)

	// echo instance
	r := router.New()

	h := handler.NewHandler(userStore, mealsStore, orderStore, restaurantStore)
	h.Register(r)

	// Server Start
	PORT := os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("PORT is empty")
		PORT = "4000"
	}

	r.Logger.Fatal(r.Start(":" + PORT))
}
