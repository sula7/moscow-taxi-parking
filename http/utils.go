package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sula7/moscow-taxi-parking/models"
)

func MakeRequest() (*models.Parkings, error) {
	url := "https://data.gov.ru/opendata/7704786030-taxiparking/data-20190906T0100.json?encoding=UTF-8"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.Status != "200" {
		body, err = ReadLocalJson("./local/data-20190906T0100.json")
		if err != nil {
			return nil, err
		}
	}

	parkings := models.Parkings{}
	err = json.Unmarshal(body, &parkings.Parking)
	if err != nil {
		return nil, err
	}

	return &parkings, nil
}

func ReadLocalJson(filePath string) ([]byte, error) {
	byteFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadAll(byteFile)
	if err != nil {
		return nil, err
	}

	return file, nil
}
