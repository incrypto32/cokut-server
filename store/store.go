package store

import (
	"github.com/incrypt0/cokut-server/brokers"
	"github.com/incrypt0/cokut-server/cokutbot"
	"github.com/incrypt0/cokut-server/common"
	"github.com/incrypt0/cokut-server/myerrors"
)

// Store is the object which abstracts db interactions
type Store struct {
	mc         string
	uc         string
	orders     string
	rc         string
	w          brokers.DBBroker
	bot        cokutbot.CokutBot
	orderCodes common.OrderCodes
	myerrors   myerrors.MyErrors
}

// NewStore creates a new store
func NewStore(mc string, uc string, orders string, rc string,

	w brokers.DBBroker,
	bot cokutbot.CokutBot,
	myerrors myerrors.MyErrors) *Store {
	return &Store{
		uc:         uc,
		rc:         rc,
		mc:         mc,
		orders:     orders,
		w:          w,
		bot:        bot,
		orderCodes: common.OrderCodes{Placed: 1, Delivered: 2, Canceled: 3},
		myerrors:   myerrors,
	}
}
