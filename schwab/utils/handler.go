package utils

import (
	"os"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"fmt"
	"strings"
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"time"

	"github.com/samjtro/go-trade/utils"
)

type TOKEN struct {
	RefreshExpiration time.Time
	Refresh string `json:"refresh_token"`
	BearerExpiration time.Time
	Bearer string `json:"access_token"`
}

// Credit: https://stackoverflow.com/questions/26916952/go-retrieve-a-string-from-between-two-characters-or-other-strings
func GetStringInBetween(str string, start string, end string) (result string) {     
	s := strings.Index(str, start)
	s += len(start)
	e := strings.Index(str[s:], end)
	e += s + e - 1

	return str[s:e] 
}

func oAuthInit() TOKEN {
	// Get Auth Code
	var (
		m sync.Mutex
	)

	m.Lock()

	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}

	resp, err := http.Get(fmt.Sprintf("https://api.schwabapi.com/v1/oauth/authorize?client_id=%s&redirect_uri=127.0.0.1", config.APPKEY))

	if err != nil {
		log.Fatalf(err.Error())
	} // WIP: else if resp.StatusCode != 404 { Handle }

	authCodeEncoded := GetStringInBetween(resp.Request.URL.String(), "?code=", "&session=")
	authCodeDecoded, err := url.QueryUnescape(authCodeEncoded)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// POST Request for Bearer/Refresh TOKENs 
	EncodedIDSecret := url.QueryEscape(fmt.Sprintf("%s:%s", config.APPKEY, config.SECRET))
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=https://example_url.com/callback_example", authCodeDecoded))))

	if err != nil {
		log.Fatalf(err.Error())
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Authorization": {EncodedIDSecret},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
	}

	res, err = res.(TOKEN)

	if err != nil {
		log.Fatalf(err.Error())
	}

	res.BearerExpiration = time.Now().Add(time.Minute * 30)
	res.RefreshExpiration = time.Now().Add(time.Hour * 168)
	
	writeOutData := fmt.Sprintf("%s,%s,%s,%s", res.RefreshExpiration, res.Refresh, res.BearerExpiration, res.Bearer)
	
	wd, err := os.Executable()

	if err != nil {
		log.Fatalf(err.Error())
	}

	wdPath := filepath.Dir(wd)

	err = ioutil.WriteFile(fmt.Sprintf("%s/db.txt", wdPath))
	m.Unlock()
	return res
}

func oAuthRefresh() string {
	var (
		m sync.Mutex
		tokens TOKEN
	)

	m.Lock()

	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}

	body, err := ioutil.ReadFile(fmt.Sprintf("%s/db.txt"), wdPath)

	if err != nil {
		log.Fatalf(err.Error())
	}

	split := strings.Split(string(body), ",")
	refreshExpiration, err := split[0].(time.Time)

	if err != nil {
		log.Fatalf(err.Error())
	}
	tokens.RefreshExpiration = refreshExpiration
	tokens.Refresh = split[1]

	bearerExpiration, err := split[0].(time.Time)

	if err != nil {
		log.Fatalf(err.Error())
	}

	tokens.BearerExpiration = bearerExpiration
	tokens.Bearer = split[3]

	EncodedIDSecret := url.QueryEscape(fmt.Sprintf("%s:%s", config.APPKEY, config.SECRET))
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", tokens.Refresh))))

	if err != nil {
		log.Fatalf(err.Error())
	}

	req.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Authorization": {EncodedIDSecret},
	}
	res, err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
	}

	tokens = res.(TOKEN)

	m.Unlock()
	return tokens.Bearer
}

// Handler is the general purpose request function for the td-ameritrade api, all functions will be routed through this handler function, which does all of the API calling work
// It performs a GET request after adding the apikey found in the .APIKEY file in the same directory as the program calling the function,
// then returns the body of the GET request's return.
// It takes one parameter:
// req = a request of type *http.Request
func Handler(req *http.Request) (string, error) {
	var (
		m sync.Mutex
		tokens TOKEN
	)

	m.Lock()

	// Going to change this to use a local godb, for now plaintext works fine
	wd, err := os.Executable()

	if err != nil {
		log.Fatalf(err.Error())
	}

	wdPath := filepath.Dir(wd)

	if _, err := os.Stat(wdPath); err == nil {
		body, err := ioutil.ReadFile(fmt.Sprintf("%s/db.txt"), wdPath)

		if err != nil {
			log.Fatalf(err.Error())
		}

		split := strings.Split(string(body), ",")
		refreshExpiration, err := split[0].(time.Time)

		if err != nil {
			log.Fatalf(err.Error())
		}
		tokens.RefreshExpiration = refreshExpiration
		tokens.Refresh = split[1]

		bearerExpiration, err := split[0].(time.Time)

		if err != nil {
			log.Fatalf(err.Error())
		}

		tokens.BearerExpiration = bearerExpiration
		tokens.Bearer = split[3]
	} else if errors.Is(err, os.ErrNotExist) {
		tokens = oAuthInit()
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	if time.Now() <= tokens.BearerExpiration {
		req.Header = http.Header{
			"Authorization": {fmt.Sprintf("Bearer %s", tokens.Bearer)},
		}
	} else {
		newBearerToken := oAuthRefresh()
		req.Header = http.Header{
			"Authorization": {fmt.Sprintf("Bearer %s", newBearerToken)},
		}
	}

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
