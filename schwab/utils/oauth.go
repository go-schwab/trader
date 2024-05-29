package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/samjtro/go-trade/utils"
	"github.com/spf13/viper"
)

func oAuthInit() TOKEN {
	var (
		m        sync.Mutex
		authCode string
	)

	m.Lock()

	// oAuth Leg 1 - App Authorization
	if _, err := os.Stat(fmt.Sprintf("%s/.foo/trade/code", utils.HomeDir())); errors.Is(err, os.ErrNotExist) {
		openBrowser(fmt.Sprintf("https://api.schwabapi.com/v1/oauth/authorize?client_id=%s&redirect_uri=%s", viper.Get("APPKEY"), viper.Get("CBURL")))
		fmt.Printf("Log into your Schwab brokerage account. Copy Error404 URL and paste it here: ")
		var urlInput string
		fmt.Scanln(&urlInput)
		authCodeEncoded := getStringInBetween(urlInput, "?code=", "&session=")
		authCode, err = url.QueryUnescape(authCodeEncoded)
		utils.Check(err)

		err = os.WriteFile(fmt.Sprintf("%s/.foo/trade/code", utils.HomeDir()), []byte(authCode), 0777)
		utils.Check(err)
	} else {
		authCodeBytes, err := os.ReadFile(fmt.Sprintf("%s/.foo/trade/code", utils.HomeDir()))
		utils.Check(err)
		authCode = string(authCodeBytes)
	}

	// oAuth Leg 2 - Access Token Creation
	authStringLegTwo := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", viper.Get("APPKEY"), viper.Get("SECRET")))))
	client := http.Client{}
	payload := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s", string(authCode), viper.Get("CBURL"))
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(payload)))
	utils.Check(err)

	req.Header = http.Header{
		"Authorization": {authStringLegTwo},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}

	res, err := client.Do(req)
	utils.Check(err)
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	utils.Check(err)
	fmt.Println(string(bodyBytes))

	tokens := parseAccessTokenResponse(string(bodyBytes))
	tokensJson, err := json.Marshal(tokens)
	utils.Check(err)

	err = os.WriteFile(fmt.Sprintf("%s/.foo/trade/bar.json", utils.HomeDir()), tokensJson, 0777)
	utils.Check(err)

	m.Unlock()
	return tokens
}

func oAuthRefresh() string {
	var (
		m sync.Mutex
	)

	m.Lock()
	tokens := readDB()

	authStringRefresh := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", viper.Get("APPKEY"), viper.Get("SECRET")))))
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", tokens.Refresh))))
	utils.Check(err)

	req.Header = http.Header{
		"Authorization": {authStringRefresh},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}

	res, err := client.Do(req)
	utils.Check(err)
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	utils.Check(err)
	fmt.Println(string(bodyBytes))

	// WIP: Working on new way to read access token response

	m.Unlock()
	return ""
}
