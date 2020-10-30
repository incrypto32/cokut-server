package utils

import (
	"encoding/json"
	"log"
	"math"
)

type Utils struct {
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2.0)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	oneighty := 180.0
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64

	la1 = lat1 * math.Pi / oneighty
	lo1 = lon1 * math.Pi / oneighty
	la2 = lat2 * math.Pi / oneighty
	lo2 = lon2 * math.Pi / oneighty

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

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
