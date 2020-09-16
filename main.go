package main

import (
	"fmt"
	"log"
	"os"

	"github.com/incrypt0/cokut-server/handler"
	"github.com/incrypt0/cokut-server/router"
	"github.com/incrypt0/cokut-server/services"
	store "github.com/incrypt0/cokut-server/stores"
)

func main() {
	// Initialize Firebase
	if _, err := services.InitFire(); err != nil {
		log.Panic(err)
	}

	// Connect to mongo
	db := services.ConnectMongo()
	users := db.Collection("users")
	meals := db.Collection("meals")
	restaurants := db.Collection("restaurants")
	orders := db.Collection("orders")

	userStore := store.NewUserStore(users)
	mealsStore := store.NewMealStore(meals, restaurants)
	restaurantStore := store.NewRestaurantStore(restaurants)
	orderStore := store.NewOrderStore(orders, restaurants)

	// echo instance
	r := router.New()
	v1 := r.Group("/api")
	h := handler.NewHandler(userStore, mealsStore, orderStore, restaurantStore)
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
