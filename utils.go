package trader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/go-schwab/oauth2ns"
	o "github.com/go-schwab/oauth2ns"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Agent struct{ client *o.AuthorizedClient }

type DB struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	Expiry       time.Time
	ExpiresIn    int64
}

var Tokens DB

func init() {
	err := godotenv.Load("*.env")
	isErrNil(err)
}

// Read in tokens from ~/.trade/bar.json
func readDB() *oauth2ns.AuthorizedClient {
	body, err := os.ReadFile(fmt.Sprintf("%s/.trade/bar.json", homeDir()))
	isErrNil(err)
	var ctx context.Context
	err = sonic.Unmarshal(body, &Tokens)
	isErrNil(err)
	token := new(oauth2.Token)
	token.AccessToken = Tokens.AccessToken
	token.RefreshToken = Tokens.RefreshToken
	token.TokenType = Tokens.TokenType
	token.Expiry = Tokens.Expiry
	token.ExpiresIn = Tokens.ExpiresIn
	c := &oauth2.Config{
		ClientID:     os.Getenv("APPKEY"),
		ClientSecret: os.Getenv("SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.schwabapi.com/v1/oauth/authorize",
			TokenURL: "https://api.schwabapi.com/v1/oauth/token",
		},
	}
	return &o.AuthorizedClient{
		c.Client(ctx, token),
		token,
	}
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

// is the err nil?
func isErrNil(err error) {
	if err != nil {
		log.Fatalf("[ERR] %s", err.Error())
	}
}

// trim one FIRST character in the string
func trimOneFirst(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1:]
}

// trim one LAST character in the string
func trimOneLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[:len(s)-1]
}

// trim one FIRST & one LAST character in the string
func trimOneFirstOneLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1 : len(s)-1]
}

// trim two FIRST & one LAST character in the string
func trimTwoFirstOneLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[2 : len(s)-1]
}

// trim one FIRST & two LAST character in the string
func trimOneFirstTwoLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1 : len(s)-2]
}

// trim one FIRST & three LAST character in the string
func trimOneFirstThreeLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1 : len(s)-3]
}

// wrapper for os.UserHomeDir()
func homeDir() string {
	dir, err := os.UserHomeDir()
	isErrNil(err)
	return dir
}

func Initiate() *Agent {
	var agent Agent
	if _, err := os.Stat(fmt.Sprintf("%s/.trade", homeDir())); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(fmt.Sprintf("%s/.trade", homeDir()), os.ModePerm)
		isErrNil(err)
		agent.client, err = o.Run()
		isErrNil(err)
		Tokens.AccessToken = agent.client.Token.AccessToken
		Tokens.RefreshToken = agent.client.Token.RefreshToken
		Tokens.TokenType = agent.client.Token.TokenType
		Tokens.Expiry = agent.client.Token.Expiry
		Tokens.ExpiresIn = agent.client.Token.ExpiresIn
		bytes, err := sonic.Marshal(Tokens)
		err = os.WriteFile(fmt.Sprintf("%s/.trade/bar.json", homeDir()), bytes, 0777)
		isErrNil(err)
	} else {
		agent.client = readDB()
		if Tokens.AccessToken == "" {
			err := os.RemoveAll(fmt.Sprintf("%s/.trade", homeDir()))
			isErrNil(err)
			log.Fatalf("[err] something went wrong - please reinitiate with 'Initiate'")
		}
	}
	return &agent
}

// Handler is the general purpose request function for the td-ameritrade api, all functions will be routed through this handler function, which does all of the API calling work
// It performs a GET request after adding the apikey found in the config.env file in the same directory as the program calling the function,
// then returns the body of the GET request's return.
// It takes one parameter:
// req = a request of type *http.Request
func (agent *Agent) Handler(req *http.Request) (*http.Response, error) {
	var err error
	if Tokens.AccessToken == "" {
		log.Fatalf("[err] no access token found, please reinitiate with 'Initiate'")
	}
	if ((&Agent{}) == agent) || ((&o.AuthorizedClient{}) == agent.client) {
		agent.client, err = o.Run()
		isErrNil(err)
		Tokens.AccessToken = agent.client.Token.AccessToken
		Tokens.RefreshToken = agent.client.Token.RefreshToken
		Tokens.TokenType = agent.client.Token.TokenType
		Tokens.Expiry = agent.client.Token.Expiry
		Tokens.ExpiresIn = agent.client.Token.ExpiresIn
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", agent.client.Token.AccessToken))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	// TODO: test this block
	var statusErr error
	switch true {
	case resp.StatusCode == 401:
		err := os.RemoveAll(fmt.Sprintf("%s/.trade", homeDir()))
		isErrNil(err)
		statusErr = errors.New("[err] invalid token - please reinitiate with 'Initiate'")
	case resp.StatusCode < 200, resp.StatusCode > 300:
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		isErrNil(err)
		statusErr = errors.New(fmt.Sprintf("[err] %d - %s", resp.StatusCode, body))
	}
	return resp, statusErr
}
