package handler

import "github.com/incrypt0/cokut-server/market"

type Handler struct {
	userStore       market.UserStore
	orderStore      market.OrderStore
	mealStore       market.MealStore
	restaurantStore market.RestaurantStore
}

func NewHandler(us market.UserStore, ms market.MealStore, os market.OrderStore, rs market.RestaurantStore) *Handler {
	return &Handler{}
}
