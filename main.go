package main

import (
	"log"

	"github.com/sula7/moscow-taxi-parking/http"
	"github.com/sula7/moscow-taxi-parking/storage"
)

func main() {
	parkings, err := http.SendRequest()
	if err != nil {
		log.Panicf("error while http GET:%v", err)
	}

	err = storage.CreateParking(parkings)
	if err != nil {
		log.Panicf("error inserting to DB:%v\n", err)
	}
}
