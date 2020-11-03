package store

import (
	"github.com/incrypt0/cokut-server/brokers"
	"github.com/incrypt0/cokut-server/common"
)

// Store is the object which abstracts db interactions
type Store struct {
	mc         string
	uc         string
	orders     string
	rc         string
	w          brokers.DBBroker
	botChannel chan string
	orderCodes common.OrderCodes
}

// NewStore creates a new store
func NewStore(mc string, uc string, orders string, rc string, w brokers.DBBroker, botChannel chan string) *Store {
	return &Store{
		uc:         uc,
		rc:         rc,
		mc:         mc,
		orders:     orders,
		w:          w,
		orderCodes: common.OrderCodes{Placed: 1, Delivered: 2, Canceled: 3},
		botChannel: botChannel,
	}
}
