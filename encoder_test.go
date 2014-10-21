package main

import (
	"./encoder"
	"log"
	"testing"
)

func TestEncode(t *testing.T) {
	//	42.446307,-76.490228

	geo, err := encoder.Encode(float64(42.446307), float64(-76.490228))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(geo)

	lat, lng := encoder.Decode(geo)
	log.Print(lat)
	log.Print(lng)

	//42.446294,-76.49014
	geo, err = encoder.Encode(float64(42.446294), float64(-76.49014))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(geo)

	lat, lng = encoder.Decode(geo)
	log.Print(lat)
	log.Print(lng)
}
