package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// create helper functions for file name search, for .APIKEY, in the project directory
// it will return a string, which is the path to the .APIKEY file in the directory
// Handler will then subsequently utilize that path for the api key element,
// thus removing the neccesity of copying around the .APIKEY file for every implementation
func keySearch() string {
	path, err := os.Getwd()
	var newPath string

	if err != nil {
		log.Fatal(err)
	}

	if path[0] == 'C' { // for Windows systems
		splitPath := strings.Split(path, "\\")

		for i := 0; i < len(splitPath)-1; i++ {
			newPath += splitPath[i] + "\\"
		}

		newPath += ".APIKEY"
	} else { // for linux/bsd systems
		splitPath := strings.Split(path, "/")

		for i := 0; i < len(splitPath)-1; i++ {
			newPath += splitPath[i] + "/"
		}

		newPath += ".APIKEY"
	}

	return newPath
}

// Handler is the general purpose request function for the td-ameritrade api
// all functions will be routed through this handler function, which does all of the API calling work
// it performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function
// it returns the body of the GET request's return
// it takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) (string, error) {
	file, err := os.Open(keySearch())

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

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
