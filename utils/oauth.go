package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func oAuthInit() TOKEN {
	var m sync.Mutex
	m.Lock()
	// Create /home/{user}/.go-trade
	err := os.Mkdir(fmt.Sprintf("%s/.trade", HomeDir()), os.ModePerm)
	Check(err)
	// oAuth Leg 1 - Authorization Code
	openBrowser(fmt.Sprintf("https://api.schwabapi.com/v1/oauth/authorize?client_id=%s&redirect_uri=%s", os.Getenv("APPKEY"), os.Getenv("CBURL")))
	fmt.Printf("Log into your Schwab brokerage account. Copy Error404 URL and paste it here: ")
	var urlInput string
	fmt.Scanln(&urlInput)
	authCodeEncoded := getStringInBetween(urlInput, "?code=", "&session=")
	authCode, err := url.QueryUnescape(authCodeEncoded)
	Check(err)
	// oAuth Leg 2 - Refresh, Bearer Tokens
	authStringLegTwo := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("APPKEY"), os.Getenv("SECRET")))))
	client := http.Client{}
	payload := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s", string(authCode), os.Getenv("CBURL"))
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(payload)))
	Check(err)
	req.Header = http.Header{
		"Authorization": {authStringLegTwo},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}
	res, err := client.Do(req)
	Check(err)
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	Check(err)
	tokens := parseAccessTokenResponse(string(bodyBytes))
	tokensJson, err := json.Marshal(tokens)
	Check(err)
	err = os.WriteFile(fmt.Sprintf("%s/.trade/bar.json", HomeDir()), tokensJson, 0777)
	Check(err)
	m.Unlock()
	return tokens
}

// Use refresh token to generate a new bearer token for authentication
func oAuthRefresh() string {
	var m sync.Mutex
	m.Lock()
	tokens := readDB()
	authStringRefresh := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("APPKEY"), os.Getenv("SECRET")))))
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", tokens.Refresh))))
	Check(err)
	req.Header = http.Header{
		"Authorization": {authStringRefresh},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}

	res, err := client.Do(req)
	Check(err)
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	Check(err)
	newTokens := parseAccessTokenResponse(string(bodyBytes))
	m.Unlock()
	return newTokens.Bearer
}
