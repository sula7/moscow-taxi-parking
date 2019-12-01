package main

import (
	v1 "github.com/sula7/moscow-taxi-parking/v1"
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

	v1.NewAPI()
}
