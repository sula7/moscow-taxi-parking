package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"

	"github.com/sula7/moscow-taxi-parking/http"
	"github.com/sula7/moscow-taxi-parking/storage"
	v1 "github.com/sula7/moscow-taxi-parking/v1"
)

type config struct {
	BindPort string `env:"BIND_PORT" envDefault:":8080"`
	DSN      string `env:"DB_CONN,required"`
	DBPwd    string `env:"DB_PWD",envDefault:""`
	FileName string `env:"FILE_NAME" envDefault:"data-20190906T0100.json"`
}

func main() {
	conf := config{}
	if err := env.Parse(&conf); err != nil {
		logrus.Fatalln("unable to get env config:", err)
	}

	parkings, err := http.GetParkingFromSource(conf.FileName)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function:": "http.GetParkingFromSource()",
			"err:":      err,
		}).Fatalln("error while http GET")
	}

	store := storage.New(conf.DSN, conf.DBPwd)

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
