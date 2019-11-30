package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendRequest() error {
	url := "https://data.gov.ru/opendata/7704786030-taxiparking/data-20190906T0100.json?encoding=UTF-8"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("response Status:\n%s", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response Body:\n%s", string(body))
	return nil
}
