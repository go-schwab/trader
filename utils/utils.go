package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type TOKEN struct {
	RefreshExpiration time.Time
	Refresh           string
	BearerExpiration  time.Time
	Bearer            string
}

func init() {
	err := LoadConfig()
	Check(err)
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
			if TrimOneFirstOneLast(x1) == "refresh_token" {
				token.Refresh = TrimOneFirstOneLast(strings.Split(x, ":")[i1+1])
			} else if TrimOneFirstOneLast(x1) == "access_token" {
				token.Bearer = TrimOneFirstOneLast(strings.Split(x, ":")[i1+1])
			}
		}
	}
	return token
}

// Read in tokens from JSON db
func readDB() TOKEN {
	var tokens TOKEN
	body, err := os.ReadFile(fmt.Sprintf("%s/.go-trade/bar.json", HomeDir()))
	Check(err)
	err = json.Unmarshal(body, &tokens)
	Check(err)
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
	Check(err)
}

func Check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func HomeDir() string {
	currentUser, err := user.Current()
	Check(err)
	return fmt.Sprintf("/home/%s", currentUser.Username)
}

func LoadConfig() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath("$HOME/.foo/trade/")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

// Trim one FIRST character in the string
func TrimOneFirst(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1:]
}

// Trim one LAST character in the string
func TrimOneLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[:len(s)-1]
}

// Trim one FIRST & one LAST character in the string
func TrimOneFirstOneLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1 : len(s)-1]
}

// Trim two FIRST & one LAST character in the string
func TrimTwoFirstOneLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[2 : len(s)-1]
}

// Trim one FIRST & two LAST character in the string
func TrimOneFirstTwoLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1 : len(s)-2]
}

// Trim one FIRST & three LAST character in the string
func TrimOneFirstThreeLast(s string) string {
	if len(s) < 1 {
		return ""
	}
	return s[1 : len(s)-3]
}

// Now returns a string; containing the current time in ISO 8601 format:
// This is the standard datetime format for quotes and dataframes across this library.
// This function uses your local `$HOME/config.env` for generation of your local time. This requires setting the UTC_DIFF variable in the following format: `-06:00`
func Now(t time.Time) string {
	return fmt.Sprintf("%d-%d-%dT%d:%d:%d%s",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		viper.Get("UTC_DIFF"))
}

func UnixToLocal(timestamp string) string {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	Check(err)
	return Now(time.Unix(i, 0))
}
