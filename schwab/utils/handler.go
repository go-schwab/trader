package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/samjtro/go-trade/utils"
)

type AccessTokenResponse struct {
	expires_in    int    `json:"expires_in"`
	token_type    string `json:"token_type"`
	scope         string `json:"scope"`
	refresh_token string `json:"refresh_token"`
	access_token  string `json:"access_token"`
	id_token      string `json:"id_token"`
}

type TOKEN struct {
	RefreshExpiration time.Time
	Refresh           string `json:"refresh_token"`
	BearerExpiration  time.Time
	Bearer            string `json:"access_token"`
}

// Credit - Old: https://stackoverflow.com/questions/26916952/go-retrieve-a-string-from-between-two-characters-or-other-strings
// Credit - New: https://go.dev/play/p/C2sZRYC15XN
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}

func oAuthInit() TOKEN {
	// Get Auth Code
	var (
		m                   sync.Mutex
		tokens              TOKEN
		accessTokenResponse AccessTokenResponse
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

	authCodeEncoded := GetStringInBetween(resp.Request.URL.String(), "?clientID=", "&region=")
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
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {EncodedIDSecret},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&accessTokenResponse)

	if err != nil {
		log.Fatalf(err.Error())
	}

	tokens.Refresh = accessTokenResponse.refresh_token
	tokens.Bearer = accessTokenResponse.access_token
	tokens.BearerExpiration = time.Now().Add(time.Minute * 30)
	tokens.RefreshExpiration = time.Now().Add(time.Hour * 168)

	writeOutData := fmt.Sprintf("%s,%s,%s,%s", utils.NowNoUTCDiff(tokens.RefreshExpiration), tokens.Refresh, utils.NowNoUTCDiff(tokens.BearerExpiration), tokens.Bearer)
	err = os.WriteFile(config.DBPATH, []byte(writeOutData), 0644)

	if err != nil {
		log.Fatalf(err.Error())
	}

	m.Unlock()
	return tokens
}

func oAuthRefresh() string {
	var (
		m                   sync.Mutex
		tokens              TOKEN
		accessTokenResponse AccessTokenResponse
	)

	m.Lock()

	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}

	body, err := os.ReadFile(config.DBPATH)

	if err != nil {
		log.Fatalf(err.Error())
	}

	split := strings.Split(string(body), ",")
	refreshExpiration, err := time.Parse("2021-08-30T08:30:00", split[0])

	if err != nil {
		log.Fatalf(err.Error())
	}

	tokens.RefreshExpiration = refreshExpiration
	tokens.Refresh = split[1]

	bearerExpiration, err := time.Parse("2021-08-30T08:30:00", split[2])

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
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {EncodedIDSecret},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
	}

	err = json.NewDecoder(res.Body).Decode(&accessTokenResponse)

	if err != nil {
		log.Fatalf(err.Error())
	}

	m.Unlock()
	return accessTokenResponse.access_token
}

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

	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Credit: https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
	if _, err := os.Stat(config.DBPATH); errors.Is(err, os.ErrNotExist) {
		tokens = oAuthInit()
	} else {
		body, err := os.ReadFile(config.DBPATH)

		if err != nil {
			log.Fatalf(err.Error())
		}

		split := strings.Split(string(body), ",")
		refreshExpiration, err := time.Parse("2021-08-30T08:30:00", split[0])

		if err != nil {
			log.Fatalf(err.Error())
		}
		tokens.RefreshExpiration = refreshExpiration
		tokens.Refresh = split[1]

		bearerExpiration, err := time.Parse("2021-08-30T08:30:00", split[2])

		if err != nil {
			log.Fatalf(err.Error())
		}

		tokens.BearerExpiration = bearerExpiration
		tokens.Bearer = split[3]
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	if !time.Now().Before(tokens.BearerExpiration) {
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
