package trader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bytedance/sonic"
	o "github.com/go-schwab/oauth2ns"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Agent struct{ client *o.AuthorizedClient }

var TOKENS oauth2.Token

func init() {
	err := godotenv.Load(findAllEnvFiles()...)
	isErrNil(err)
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

// wrapper for os.UserHomeDir()
func homeDir() string {
	dir, err := os.UserHomeDir()
	isErrNil(err)
	return dir
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

// Read in tokens from ~/.trade/bar.json
func readDB() *o.AuthorizedClient {
	body, err := os.ReadFile(fmt.Sprintf("%s/.trade/bar.json", homeDir()))
	isErrNil(err)
	var ctx context.Context
	err = sonic.Unmarshal(body, &TOKENS)
	isErrNil(err)
	c := &oauth2.Config{
		ClientID:     os.Getenv("APPKEY"),
		ClientSecret: os.Getenv("SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.schwabapi.com/v1/oauth/authorize",
			TokenURL: "https://api.schwabapi.com/v1/oauth/token",
		},
	}
	return &o.AuthorizedClient{
		c.Client(ctx, &TOKENS),
		&TOKENS,
	}
}

func Initiate() *Agent {
	var agent Agent
	// TODO: test this block, this is to attempt to resolve the error described in #67
	if _, err := os.Stat(fmt.Sprintf("%s/.trade/bar.json", homeDir())); errors.Is(err, os.ErrNotExist) {
		if _, err := os.Stat(fmt.Sprintf("%s/.trade", homeDir())); !errors.Is(err, os.ErrNotExist) {
			err = os.RemoveAll(fmt.Sprintf("%s/.trade", homeDir()))
			isErrNil(err)
		}
		err = os.Mkdir(fmt.Sprintf("%s/.trade", homeDir()), os.ModePerm)
		isErrNil(err)
		agent.client, err = o.Initiate(os.Getenv("APPKEY"), os.Getenv("SECRET"), os.Getenv("CBURL"))
		isErrNil(err)
		TOKENS = *agent.client.Token
		bytes, err := sonic.Marshal(TOKENS)
		err = os.WriteFile(fmt.Sprintf("%s/.trade/bar.json", homeDir()), bytes, 0777)
		isErrNil(err)
	} else {
		agent.client = readDB()
		if TOKENS.AccessToken == "" {
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
	if TOKENS.AccessToken == "" {
		log.Fatalf("[err] no access token found, please reinitiate with 'Initiate'")
	}
	if ((&Agent{}) == agent) || ((&o.AuthorizedClient{}) == agent.client) {
		// TODO: this theoretically works but results in an error for the oauth implementation
		// going to do some testing now, but pushing as it is in what is a theoretically working state
		agent.client, err = o.Initiate(os.Getenv("APPKEY"), os.Getenv("SECRET"), os.Getenv("CBURL"))
		isErrNil(err)
		TOKENS = *agent.client.Token
	}
	resp, err := agent.client.Do(req)
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
