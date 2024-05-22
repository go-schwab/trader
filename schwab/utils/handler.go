package utils

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"fmt"
	"strings"
	"bytes"
	"encoding/base64"
)

type Token struct {
	refresh string `json:"refresh_token"` 
	access string `json:"access_token"`
}

func init() {
	config, err := LoadConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}
}

// Credit: https://stackoverflow.com/questions/26916952/go-retrieve-a-string-from-between-two-characters-or-other-strings
func GetStringInBetween(str string, start string, end string) (result string) {     
	s := strings.Index(str, start)
	s += len(start)
	e := strings.Index(str[s:], end)
	e += s + e - 1

	return str[s:e] 
}

func oAuth() {
	// Get Auth Code
	var (
		m sync.Mutex
	)

	resp, err := http.Get(fmt.Sprintf("https://api.schwabapi.com/v1/oauth/authorize?client_id=%s&redirect_uri=127.0.0.1", config.APPKEY))

	if err != nil {
		log.Fatalf(err.Error())
	} else if resp.StatusCode != 404 {
		log.Fatalf(err.Error())
	}

	authCodeEncoded := GetStringInBetween(resp.Request.URL.String(), "?code=", "&session=")

	// POST Request for Bearer/Refresh Tokens
	EncodedIDSecret := base64.URLEncoding.EncodeToString(fmt.Sprintf("%s:%s", config.APPKEY, config.SECRET)
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer(fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=https://example_url.com/callback_example"), url.QueryUnescape(authCodeEncoded)))
	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Authorization": {EncodedIDSecret},
	}
	res, err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
	}

	res = res.(Tokens)
}

// Handler is the general purpose request function for the td-ameritrade api, all functions will be routed through this handler function, which does all of the API calling work
// It performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function,
// then returns the body of the GET request's return.
// It takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) (string, error) {
	var (
		m sync.Mutex
	)

	m.Lock()

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	errorCode := resp.StatusCode
	bodyBytes, err := io.ReadAll(resp.Body)
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
