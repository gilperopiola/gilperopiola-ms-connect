package connect

import (
	"io/ioutil"
	"net/http"
)

func sendHTTPRequest(req *http.Request) (response string) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if b, err := ioutil.ReadAll(resp.Body); err == nil {
		return string(b)
	}

	return ""
}
