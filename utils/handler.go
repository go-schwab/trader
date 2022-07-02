package utils

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
)

// create helper functions for file name search, for .APIKEY, in the project directory
// it will return a string, which is the path to the .APIKEY file in the directory
// Handler will then subsequently utilize that path for the api key element,
// thus removing the neccesity of copying around the .APIKEY file for every implementation
// func keySeach() string {}

// Handler is the general purpose request function for the td-ameritrade api
// all functions will be routed through this handler function, which does all of the API calling work
// it performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function
// it returns the body of the GET request's return
// it takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) string {
	file, _ := os.Open(".APIKEY")
	s := bufio.NewScanner(file)
	var APIKEY string

	for s.Scan() {
		APIKEY += s.Text()
	}

	q := req.URL.Query()
	q.Add("apikey", APIKEY)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return string(body)
}
