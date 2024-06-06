package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("config.env")
	Check(err)
}

type TOKEN struct {
	RefreshExpiration time.Time
	Refresh           string
	BearerExpiration  time.Time
	Bearer            string
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
	body, err := os.ReadFile(fmt.Sprintf("%s/.trade/bar.json", HomeDir()))
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
		log.Fatalf("Unsupported platform.")
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
	var homedir string
	switch runtime.GOOS {
	case "linux":
		homedir = "/home/" + currentUser.Username
	case "windows":
		homedir = "C:\\Users\\" + currentUser.Username
	case "darwin":
		homedir = "/users/" + currentUser.Username
	default:
		log.Fatalf("Unsupported platform.")
	}
	return homedir
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
