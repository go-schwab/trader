package trader

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	o "github.com/go-schwab/utils/oauth"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Agent struct {
	Client *o.AuthorizedClient
	Tokens Token
	Linux  bool
}

type Token struct {
	RefreshExpiration time.Time
	Refresh           string
	BearerExpiration  time.Time
	Bearer            string
}

var (
	APPKEY string
	SECRET string
	CBURL  string
)

func init() {
	err := godotenv.Load(findAllEnvFiles()...)
	isErrNil(err)
	APPKEY = os.Getenv("APPKEY")
	SECRET = os.Getenv("SECRET")
	CBURL = os.Getenv("CBURL")
}

// is the err nil?
func isErrNil(err error) {
	if err != nil {
		log.Fatalf("[fatal] %s", err.Error())
	}
}

// trim one FIRST & one LAST character in the string
func trimOneFirstOneLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1 : len(s)-1]
}

// Helper: parse access token response
func parseAccessTokenResponse(s string) Token {
	token := Token{
		RefreshExpiration: time.Now().Add(time.Hour * 168),
		BearerExpiration:  time.Now().Add(time.Minute * 30),
	}
	for _, x := range strings.Split(s, ",") {
		for i1, x1 := range strings.Split(x, ":") {
			if trimOneFirstOneLast(x1) == "refresh_token" {
				token.Refresh = trimOneFirstOneLast(strings.Split(x, ":")[i1+1])
			} else if trimOneFirstOneLast(x1) == "access_token" {
				token.Bearer = trimOneFirstOneLast(strings.Split(x, ":")[i1+1])
			}
		}
	}
	return token
}

// find all env files
func findAllEnvFiles() []string {
	var files []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		split := strings.Split(d.Name(), ".")
		if len(split) > 1 {
			if split[1] == "env" {
				files = append(files, d.Name())
			}
		}
		return err
	})
	isErrNil(err)
	return files
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
		log.Fatalf("Unsupported platform.")
	}
	isErrNil(err)
}

// Execute a command @ stdin, receive stdout
func execCommand(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Fatalf(err.Error())
	}
}

// Read in tokens from .json
func readLinuxDB() Token {
	var tokens Token
	body, err := os.ReadFile(".json")
	isErrNil(err)
	err = json.Unmarshal(body, &tokens)
	isErrNil(err)
	return tokens
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

// Read in tokens from .json
func readDB() Agent {
	var tok *oauth2.Token
	body, err := os.ReadFile(".json")
	isErrNil(err)
	err = sonic.Unmarshal(body, &tok)
	isErrNil(err)
	conf := &oauth2.Config{
		ClientID:     APPKEY, // Schwab App Key
		ClientSecret: SECRET, // Schwab App Secret
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.schwabapi.com/v1/oauth/authorize",
			TokenURL: "https://api.schwabapi.com/v1/oauth/token",
		},
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	sslcli := &http.Client{Transport: tr}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, sslcli)
	return Agent{
		Client: &o.AuthorizedClient{
			conf.Client(ctx, tok),
			tok,
		},
	}
}

// FOR LINUX USERS
func InitiateLinux() *Agent {
	var agent Agent
	if _, err := os.Stat(".json"); errors.Is(err, os.ErrNotExist) {
		// oAuth Leg 1 - Authorization Code
		openBrowser(fmt.Sprintf("https://api.schwabapi.com/v1/oauth/authorize?client_id=%s&redirect_uri=%s", os.Getenv("APPKEY"), os.Getenv("CBURL")))
		fmt.Printf("Log into your Schwab brokerage account. Copy Error404 URL and paste it here: ")
		var urlInput string
		fmt.Scanln(&urlInput)
		authCodeEncoded := getStringInBetween(urlInput, "?code=", "&session=")
		authCode, err := url.QueryUnescape(authCodeEncoded)
		isErrNil(err)
		// oAuth Leg 2 - Refresh, Bearer Tokens
		authStringLegTwo := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("APPKEY"), os.Getenv("SECRET")))))
		client := http.Client{}
		payload := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s", string(authCode), os.Getenv("CBURL"))
		req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(payload)))
		isErrNil(err)
		req.Header = http.Header{
			"Authorization": {authStringLegTwo},
			"Content-Type":  {"application/x-www-form-urlencoded"},
		}
		res, err := client.Do(req)
		isErrNil(err)
		defer res.Body.Close()
		bodyBytes, err := io.ReadAll(res.Body)
		isErrNil(err)
		agent.Tokens = parseAccessTokenResponse(string(bodyBytes))
		bytes, err := sonic.Marshal(agent.Tokens)
		isErrNil(err)
		err = os.WriteFile(".json", bytes, 0777)
		isErrNil(err)
	} else {
		agent.Tokens = readLinuxDB()
		if agent.Tokens.Bearer == "" {
			err := os.Remove(".json")
			isErrNil(err)
			log.Fatalf("[err] please reinitiate, something went wrong\n")
		}
	}
	agent.Linux = true
	return &agent
}

// FOR MAC, WINDOWS USERS
func Initiate() *Agent {
	var agent Agent
	if _, err := os.Stat(".json"); errors.Is(err, os.ErrNotExist) {
		//execCommand("openssl req -x509 -out localhost.crt -keyout localhost.key   -newkey rsa:2048 -nodes -sha256   -subj '/CN=localhost' -extensions EXT -config <(;printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS.1:localhost,IP:127.0.0.1\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")")
		agent = Agent{Client: o.Initiate(APPKEY, SECRET)}
		bytes, err := sonic.Marshal(agent.Client.Token)
		isErrNil(err)
		err = os.WriteFile(".json", bytes, 0777)
		isErrNil(err)
	} else {
		agent = readDB()
	}
	agent.Linux = false
	return &agent
}

// Use refresh token to generate a new bearer token for authentication
func (agent *Agent) refresh() {
	oldTokens := readLinuxDB()
	authStringRefresh := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("APPKEY"), os.Getenv("SECRET")))))
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.schwabapi.com/v1/oauth/token", bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", oldTokens.Refresh))))
	isErrNil(err)
	req.Header = http.Header{
		"Authorization": {authStringRefresh},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}
	res, err := client.Do(req)
	isErrNil(err)
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	isErrNil(err)
	agent.Tokens = parseAccessTokenResponse(string(bodyBytes))
}

// Handler is the general purpose request function for the td-ameritrade api, all functions will be routed through this handler function, which does all of the API calling work
// It performs a GET request after adding the apikey found in the config.env file in the same directory as the program calling the function,
// then returns the body of the GET request's return.
// It takes one parameter:
// req = a request of type *http.Request
func (agent *Agent) Handler(req *http.Request) (*http.Response, error) {
	if agent.Linux {
		if (&Agent{}) == agent {
			log.Fatal("[fatal] empty agent - call 'Agent.Initiate' before making any API function calls.")
			// TODO: auto reinitiate?
		}
		if !time.Now().Before(agent.Tokens.BearerExpiration) {
			agent.refresh()
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", agent.Tokens.Bearer))
		client := http.Client{}
		resp, err := client.Do(req)
		isErrNil(err)
		switch true {
		case resp.StatusCode == 200:
			return resp, nil
		case resp.StatusCode == 401:
			log.Fatal("[fatal] invalid token - please reinitiate with 'Initiate'")
			// TODO: auto reinitiate?
		case resp.StatusCode < 200, resp.StatusCode > 300:
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			isErrNil(err)
			log.Println("[err] ", string(body))
		}
	} else {
		if agent.Client.Token.AccessToken == "" {
			log.Fatal("[fatal] no access token found, please reinitiate with 'Initiate'")
			// TODO: auto reinitiate?
		}
		resp, err := agent.Client.Do(req)
		if err != nil {
			return resp, err
		}
		switch true {
		case resp.StatusCode == 200:
			return resp, nil
		case resp.StatusCode == 401:
			log.Fatal("[fatal] invalid token - please reinitiate with 'Initiate'")
			// TODO: auto reinitiate?
		case resp.StatusCode < 200, resp.StatusCode > 300:
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			isErrNil(err)
			log.Println("[err] ", string(body))
		}
	}
	return nil, nil
}
