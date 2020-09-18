package handler2

import (
	"github.com/incrypt0/cokut-server/store"
)

type Handler struct {
	store *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{
		store: store,
	}
}
