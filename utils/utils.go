package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

// Print a model
func PrintModel(u interface{}) string {
	fmt.Println("\n________Print Model_______")
	fmt.Println()
	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	s := string(b)
	return s
}
