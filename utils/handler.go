package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Handler is the general purpose request function for the td-ameritrade api, all functions will be routed through this handler function, which does all of the API calling work
// It performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function,
// then returns the body of the GET request's return.
// It takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) (string, error) {
	keyPath := "~/.APIKEY"
	file, err := os.Open(keyPath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	var APIKEY string
	s := bufio.NewScanner(file)

	for s.Scan() {
		APIKEY += s.Text()
	}

	q := req.URL.Query()
	q.Add("apikey", APIKEY)
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		errorCode := resp.StatusCode

		switch errorCode {
		case 500:
			log.Fatalf("Error %d - Invalid Authentication; Check your API Key", errorCode)
		case 400:
			log.Fatalf("Error %d - Bad Request; Ensure your parameters are correct", errorCode)
		}

	}

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
