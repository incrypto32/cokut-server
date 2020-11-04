package handler2

import (
	"github.com/incrypt0/cokut-server/brokers"
	"github.com/incrypt0/cokut-server/myerrors"
	"github.com/labstack/echo/v4"
)

// Handler is main api handler object which controls every network interactions.
type Handler struct {
	store         brokers.CokutBroker
	fireAuthMWare echo.MiddlewareFunc
	adminChecker  echo.MiddlewareFunc
	Domain        string
	myerrors      myerrors.MyErrors
}

type ManyResultFunc func() ([]interface{}, error)

type FilteredManyResultFunc func(string) ([]interface{}, error)

// NewHandler is creates a new Handler object.
func NewHandler(store brokers.CokutBroker,
	fireAuthMWare echo.MiddlewareFunc,
	adminChecker echo.MiddlewareFunc,
	myerrors myerrors.MyErrors) *Handler {
	return &Handler{
		store:         store,
		fireAuthMWare: fireAuthMWare,
		adminChecker:  adminChecker,
		myerrors:      myerrors,
	}
}
