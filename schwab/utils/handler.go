package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/samjtro/go-trade/utils"
)

// Handler is the general purpose request function for the td-ameritrade api, all functions will be routed through this handler function, which does all of the API calling work
// It performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function,
// then returns the body of the GET request's return.
// It takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) (string, error) {
	var (
		m      sync.Mutex
		tokens TOKEN
	)
	m.Lock()
	// Check if program has been run before by verifying the existence of /home/{user}/.go-trade
	if _, err := os.Stat(fmt.Sprintf("%s/.go-trade", utils.HomeDir())); errors.Is(err, os.ErrNotExist) {
		tokens = oAuthInit()
	} else {
		tokens = readDB()
	}
	// Check if bearer token is still valid
	if !time.Now().Before(tokens.BearerExpiration) {
		tokens.Bearer = oAuthRefresh()
	}
	// Craft, send request
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokens.Bearer))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Grab return body
	errorCode := resp.StatusCode
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)
	if errorCode < 200 || errorCode > 300 {
		log.Fatalf("Error %d - %s", errorCode, body)
	}
	m.Unlock()
	return body, nil
}
