package utils

import (
	"os"
	"bufio"
	"net/http"
	"io/ioutil"
)

// handler is the general purpose request function for the td-ameritrade api
// all functions will be routed through this handler function, which does all of the API calling work
func handler(req *http.Request) string {
	file,_ := os.Open(".APIKEY")
	s := bufio.NewScanner(file)
	var APIKEY string
	
	for s.Scan() {
		APIKEY += s.Text()
	}

	q := req.URL.Query()
	q.Add("apikey",APIKEY)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp,_ := client.Do(req)
	body,_ := ioutil.ReadAll(resp.Body)
	
	resp.Body.Close()

	return string(body)
}
