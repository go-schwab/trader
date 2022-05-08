package handler

import (
	"os"
	"bufio"
	"net/http"
	"io/ioutil"
)

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
