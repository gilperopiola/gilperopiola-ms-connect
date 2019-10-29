package connect

import (
	"bytes"
	"net/http"
)

func SendMail(to, subject, text, html string) (response string) {
	msMailingURL := "localhost:8067/v1"
	endpointURL := msMailingURL + "/SendMail"

	httpRequestBody := `{
		"to": "` + to + `",
		"subject": "` + subject + `",
		"text": "` + text + `",
		"html": "` + html + `"
	}`

	request, _ := http.NewRequest("POST", endpointURL, bytes.NewReader([]byte(httpRequestBody)))
	return sendHTTPRequest(request)
}

type Email struct {
	To      string
	Subject string
	Text    string
	HTML    string
}
