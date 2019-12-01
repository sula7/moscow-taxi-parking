package main

import (
	"log"

	"github.com/sula7/moscow-taxi-parking/http"
	"github.com/sula7/moscow-taxi-parking/storage"
)

func main() {
	parkings, err := http.SendRequest()
	if err != nil {
		log.Fatalln("error while http GET: ", err)
	}

	err = storage.Migrate()
	if err != nil {
		log.Fatalln("Unable to run DB migrations: ", err)
	}

	err = storage.CreateParking(parkings)
	if err != nil {
		log.Fatalln("error inserting to DB: ", err)
	}
}
