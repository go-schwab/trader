package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// create helper functions for file name search, for .APIKEY, in the project directory
// it will return a string, which is the path to the .APIKEY file in the directory
// Handler will then subsequently utilize that path for the api key element,
// thus removing the neccesity of copying around the .APIKEY file for every implementation
func keySearch() (string, error) {
	path, err := os.Getwd()

	if err != nil {
		return "", err
	}

	var newPath string

	if path[0] == 'C' { // for Windows systems
		splitPath := strings.Split(path, "\\")

		for i := 0; i < len(splitPath)-1; i++ {
			newPath += splitPath[i] + "\\"
		}
	} else { // for linux/bsd systems
		splitPath := strings.Split(path, "/")

		for i := 0; i < len(splitPath)-1; i++ {
			newPath += splitPath[i] + "/"
		}

	}

	newPath += ".APIKEY"

	return newPath, nil
}

// Handler is the general purpose request function for the td-ameritrade api
// all functions will be routed through this handler function, which does all of the API calling work
// it performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function
// it returns the body of the GET request's return
// it takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) (string, error) {
	keyPath, err := keySearch()

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

	if resp.StatusCode <= 200 || resp.StatusCode >= 300 {
		fmt.Println(fmt.Sprintf("Error %d", resp.StatusCode))
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
