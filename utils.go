package trader

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bytedance/sonic"
	o "github.com/go-schwab/utils/oauth"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Agent struct{ *o.AuthorizedClient }

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
		&o.AuthorizedClient{
			conf.Client(ctx, tok),
			tok,
		},
	}
}

func Initiate() *Agent {
	var agent Agent
	// TODO: test this block, this is to attempt to resolve the error described in #67
	if _, err := os.Stat(".json"); errors.Is(err, os.ErrNotExist) {
		agent = Agent{o.Initiate(APPKEY, SECRET)}
		bytes, err := sonic.Marshal(agent.Token)
		isErrNil(err)
		err = os.WriteFile(".json", bytes, 0777)
		isErrNil(err)
	} else {
		agent = readDB()
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
	if agent.Token.AccessToken == "" {
		log.Fatal("[fatal] no access token found, please reinitiate with 'Initiate'")
		// TODO: auto reinitiate?
	}
	resp, err := agent.Do(req)
	if err != nil {
		return resp, err
	}
	switch true {
	case resp.StatusCode == 401:
		log.Fatal("[fatal] invalid token - please reinitiate with 'Initiate'")
	case resp.StatusCode < 200, resp.StatusCode > 300:
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		isErrNil(err)
		log.Println("[err] ", string(body))
	}
	return resp, nil
}
