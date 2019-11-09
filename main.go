package connect

import (
	"io/ioutil"
	"net/http"
)

func sendHTTPRequest(req *http.Request) (status int, response string) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, err.Error()
	}
	defer resp.Body.Close()

	if b, err := ioutil.ReadAll(resp.Body); err == nil {
		return resp.StatusCode, string(b)
	}

	return 400, err.Error()
}
