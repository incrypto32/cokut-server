package store

import "github.com/incrypt0/cokut-server/workers"

type Store struct {
	mc     string
	uc     string
	orders string
	rc     string
	w      *workers.Worker
}

func NewStore(mc string, uc string, orders string, rc string, w *workers.Worker) *Store {
	return &Store{
		uc:     uc,
		rc:     rc,
		mc:     mc,
		orders: orders,
		w:      w,
	}
}
