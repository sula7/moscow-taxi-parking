package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sula7/moscow-taxi-parking/models"
)

// GetParkingFromSource makes HTTP request to url and parses data Parking struct.
// If URL is unavailable then reads local file
func GetParkingFromSource(url, fileName string) (parkings models.Parkings, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.Status != "200" {
		body, err = ReadLocalJson(fileName)
		if err != nil {
			return
		}
	}

	err = json.Unmarshal(body, &parkings.Parking)
	if err != nil {
		return
	}

	return parkings, nil
}

// ReadLocalJson reads local .json and parses it to []byte
func ReadLocalJson(fileName string) ([]byte, error) {
	byteFile, err := os.Open("./local/" + fileName)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadAll(byteFile)
	if err != nil {
		return nil, err
	}

	return file, nil
}
