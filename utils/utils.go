package utils

import (
	"encoding/json"
	"log"
)

// Print a model
func PrintModel(u interface{}) string {
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
