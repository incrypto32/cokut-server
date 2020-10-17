package utils

import (
	"encoding/json"
	"log"
)

// ModelToString .
func ModelToString(u interface{}) string {
	// Enthenkilum preshnam vannal mone orku ni print cheyypikkanam ivide print cheyyilla
	log.Println("\n________Print Model_______")
	log.Println()

	b, err := json.MarshalIndent(u, "", "  ")

	if err != nil {
		log.Println(err)
		return err.Error()
	}

	s := string(b)

	return s
}

// NewBool
func NewBool(val bool) *bool {
	b := val
	return &b
}
