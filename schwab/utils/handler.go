package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/samjtro/go-trade/utils"
)

type AccessTokenResponse struct {
	expires_in    int
	token_type    string
	scope         string
	refresh_token string
	access_token  string
	id_token      string
}

type TOKEN struct {
	RefreshExpiration time.Time
	Refresh           string
	BearerExpiration  time.Time
	Bearer            string
}

// Credit: https://go.dev/play/p/C2sZRYC15XN
func getStringInBetween(str string, start string, end string) (result string) {
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

// Credit: https://gist.github.com/hyg/9c4afcd91fe24316cbf0
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func readDB() TOKEN {
	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}

	body, err := os.ReadFile(config.DBPATH)

	if err != nil {
		log.Fatalf(err.Error())
	}

	split := strings.Split(string(body), ",")
	refreshAsInt, err := strconv.ParseInt(split[0], 10, 64)

	if err != nil {
		log.Fatalf(err.Error())
	}

	bearerAsInt, err := strconv.ParseInt(split[2], 10, 64)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return TOKEN{
		RefreshExpiration: time.Unix(refreshAsInt, 0),
		Refresh:           split[1],
		BearerExpiration:  time.Unix(bearerAsInt, 0),
		Bearer:            split[3],
	}
}

func oAuthInit() TOKEN {
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

	// oAuth Leg 1 - App Authorization
	openBrowser(fmt.Sprintf("https://api.schwabapi.com/v1/oauth/authorize?client_id=%s&redirect_uri=%s", config.APPKEY, config.CBURL))
	fmt.Printf("Log into your Schwab brokerage account. Copy Error404 URL and paste it here: ")
	var urlInput string
	fmt.Scanln(&urlInput)
	authCodeEncoded := getStringInBetween(urlInput, "?code=", "&session=")
	authCodeDecoded, err := url.QueryUnescape(authCodeEncoded)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// oAuth Leg 2 - Access Token Creation
	authStringLegTwo := fmt.Sprintf("Basic %s:%s", base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(config.APPKEY))), config.SECRET)
	fmt.Println(authStringLegTwo)
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s", authCodeDecoded, config.CBURL))))

	if err != nil {
		log.Fatalf(err.Error())
	}

	req.Header = http.Header{
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {authStringLegTwo},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
	}

	defer res.Body.Close()

	// Credit: https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go
	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(bodyBytes))

	tokens.Refresh = accessTokenResponse.refresh_token
	tokens.Bearer = accessTokenResponse.access_token
	tokens.BearerExpiration = time.Now().Add(time.Minute * 30)
	tokens.RefreshExpiration = time.Now().Add(time.Hour * 168)

	writeOutData := fmt.Sprintf("%d,%s,%d,%s", tokens.RefreshExpiration.Unix(), tokens.Refresh, tokens.BearerExpiration.Unix(), tokens.Bearer)
	err = os.WriteFile(config.DBPATH, []byte(writeOutData), 0755)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Done.")
	m.Unlock()
	return tokens
}

func oAuthRefresh() string {
	var (
		m                   sync.Mutex
		accessTokenResponse AccessTokenResponse
	)

	m.Lock()
	tokens := readDB()
	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}

	// POST Request
	EncodedIDSecret := base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(fmt.Sprintf("Basic %s:%s", config.APPKEY, config.SECRET))))
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

	// Credit: https://golangtutorial.dev/tips/check-if-a-file-exists-or-not-in-go/
	// Check if DBPATH exists; if it does, read in from file; if it does not, execuate oAuthInit()
	if _, err := os.Stat(config.DBPATH); errors.Is(err, os.ErrNotExist) {
		tokens = oAuthInit()
	} else {
		tokens = readDB()
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
	fmt.Println(body)

	if err != nil {
		return "", err
	}

	if errorCode < 200 || errorCode > 300 {
		log.Fatalf("Error %d - %s", errorCode, body)
	}

	m.Unlock()
	return body, nil
}
