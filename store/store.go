package store

import (
	"github.com/incrypt0/cokut-server/brokers"
)

type Store struct {
	mc     string
	uc     string
	orders string
	rc     string
	w      brokers.DbBroker
}

func NewStore(mc string, uc string, orders string, rc string, w brokers.DbBroker) *Store {
	return &Store{
		uc:     uc,
		rc:     rc,
		mc:     mc,
		orders: orders,
		w:      w,
	}
}
