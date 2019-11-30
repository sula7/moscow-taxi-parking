package main

import (
	"github.com/sula7/moscow-taxi-parking/http"
	"log"
)

func main() {
	err := http.SendRequest()
	if err != nil {
		log.Panicf("error while http GET:%v", err)
	}
}
