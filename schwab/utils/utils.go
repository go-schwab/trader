package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/samjtro/go-trade/utils"
)

type TOKEN struct {
	RefreshExpiration time.Time
	Refresh           string
	BearerExpiration  time.Time
	Bearer            string
}

func init() {
	err := utils.LoadConfig()
	utils.Check(err)
}

/*func trimFirstJSONElement(s string) string {
	return s[2 : len(s)-1]
}

func trimLastJSONElement(s string) string {
	return s[1 : len(s)-2]
}*/

// oAuthInit() helper func, parse access token response
func parseAccessTokenResponse(s string) TOKEN {
	token := TOKEN{
		RefreshExpiration: time.Now().Add(time.Hour * 168),
		BearerExpiration:  time.Now().Add(time.Minute * 30),
	}

	for _, x := range strings.Split(s, ",") {
		for i1, x1 := range strings.Split(x, ":") {
			if utils.TrimOneFirstOneLast(x1) == "refresh_token" {
				token.Refresh = utils.TrimOneFirstOneLast(strings.Split(x, ":")[i1+1])
			} else if utils.TrimOneFirstOneLast(x1) == "access_token" {
				token.Bearer = utils.TrimOneFirstOneLast(strings.Split(x, ":")[i1+1])
			}
		}
	}

	return token
}

// Read in tokens from JSON db
func readDB() TOKEN {
	var tokens TOKEN
	body, err := os.ReadFile(fmt.Sprintf("%s/.go-trade/bar.json", utils.HomeDir()))
	utils.Check(err)

	err = json.Unmarshal(body, &tokens)
	utils.Check(err)

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

	utils.Check(err)
}
