package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

// Handler is the general purpose request function for the td-ameritrade api, all functions will be routed through this handler function, which does all of the API calling work
// It performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function,
// then returns the body of the GET request's return.
// It takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) (string, error) {
	var (
		m      sync.Mutex
		APIKEY string
	)

	dir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	switch dir[0] {
	case 'C':
		dir += "\\.APIKEY"
	default:
		dir += "/.APIKEY"
	}

	m.Lock()

	file, err := os.Open(dir)

	if err != nil {
		return "", err
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		APIKEY += s.Text()
	}

	m.Unlock()

	q := req.URL.Query()
	q.Add("apikey", APIKEY)
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	m.Lock()
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

	m.Unlock()
	return body, nil
}
