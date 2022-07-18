package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// KeySearch returns the parent directory + ".APIKEY".
// This will later serve as a search function for the entire home directory, hence the name, but I don't want to spend that much time on this yet when this accomplishes the same thing.
// Namely, the ability for .APIKEY to remain outside of the project folder
func KeySearch() (string, error) {
	path, err := os.Getwd()

	if err != nil {
		return "", err
	}

	var newPath string
	var length int

	if path[0] == 'C' { // for Windows systems
		splitPath := strings.Split(path, "\\")

		for i, x := range splitPath {
			if x == "Users" {
				length += i + 2
			}
		}

		for i := 0; i < length; i++ {
			newPath += splitPath[i] + "\\"
		}
	} else { // for linux/bsd/mac systems
		splitPath := strings.Split(path, "/")

		for i, x := range splitPath {
			if x == "home" || x == "Users" {
				length += i + 2
			}
		}

		for i := 0; i < length; i++ {
			newPath += splitPath[i] + "/"
		}

	}

	newPath += ".APIKEY"

	return newPath, nil
}

// Handler is the general purpose request function for the td-ameritrade api, all functions will be routed through this handler function, which does all of the API calling work
// It performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function,
// then returns the body of the GET request's return.
// It takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) (string, error) {
	keyPath, err := KeySearch()

	if err != nil {
		return "", err
	}

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

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	errorCode := resp.StatusCode
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)

	if err != nil {
		return "", err
	}

	if errorCode < 200 || errorCode > 300 {
		log.Fatalf("Error %d - %s", errorCode, body)
	}

	return body, nil
}
