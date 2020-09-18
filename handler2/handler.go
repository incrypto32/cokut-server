package handler2

import (
	"github.com/incrypt0/cokut-server/brokers"
	"github.com/incrypt0/cokut-server/store"
)

type Handler struct {
	store brokers.CokutBroker
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{
		store: store,
	}
}
