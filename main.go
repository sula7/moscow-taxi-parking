package main

import (
	"github.com/sirupsen/logrus"

	"github.com/sula7/moscow-taxi-parking/config"
	"github.com/sula7/moscow-taxi-parking/http"
	"github.com/sula7/moscow-taxi-parking/storage"
	v1 "github.com/sula7/moscow-taxi-parking/v1"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "config.Get()",
			"err:":      err,
		}).Fatalln("unable to get env config:", err)
	}

	parkings, err := http.GetParkingFromSource(conf.FileName)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "http.GetParkingFromSource()",
			"err:":      err,
		}).Fatalln("error while http GET")
	}

	store, err := storage.New(conf.DSN, conf.DBPwd)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "http.storage.New()",
			"err:":      err,
		}).Fatalln("error establishing connection to DB")
	}

	defer store.Close()

	err = store.CreateParkings(parkings)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "store.CreateParkings()",
			"err:":      err,
		}).Fatalln("error creating parkings info in DB")
	}

	v1.NewAPI(store, conf.BindPort)
}
