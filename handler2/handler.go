package handler2

import (
	"github.com/incrypt0/cokut-server/brokers"
	"github.com/incrypt0/cokut-server/store"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store         brokers.CokutBroker
	fireAuthMWare echo.MiddlewareFunc
}

func NewHandler(store *store.Store, fireAuthMWare echo.MiddlewareFunc) *Handler {
	return &Handler{
		store:         store,
		fireAuthMWare: fireAuthMWare,
	}
}
