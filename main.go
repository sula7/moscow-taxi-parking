package main

import (
	"github.com/sirupsen/logrus"

	"github.com/sula7/moscow-taxi-parking/http"
	"github.com/sula7/moscow-taxi-parking/storage"
	v1 "github.com/sula7/moscow-taxi-parking/v1"
)

func main() {
	parkings, err := http.MakeRequest()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "http.MakeRequest()",
			"err:":      err,
		}).Fatalln("error while http GET")
	}

	store, err := storage.New("root:root@tcp(localhost:3306)/parkings")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "storage.New()",
			"err:":      err,
		}).Fatalln("error creating new storage")
	}

	defer store.Close()

	err = storage.Migrate("mysql://root:root@tcp(localhost:3306)/parkings")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "storage.Migrate()",
			"err:":      err,
		}).Fatalln("Unable to run DB migrations")
	}

	err = store.CreateParkings(parkings)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "store.CreateParkings()",
			"err:":      err,
		}).Fatalln("error creating parkings info in DB")
	}

	v1.NewAPI(store)
}
