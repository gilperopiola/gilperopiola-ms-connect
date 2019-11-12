package connect

import "github.com/gilperopiola/frutils"

const msMailingURL = "http://localhost:9004/v1"

func SendMail(to, subject, text, html string) (status int, response string) {
	endpointURL := msMailingURL + "/SendMail"

	httpRequestBody := `{
		"to": "` + to + `",
		"subject": "` + subject + `",
		"text": "` + text + `",
		"html": "` + html + `"
	}`

	return frutils.SendHTTPRequest("POST", endpointURL, httpRequestBody)
}

type Email struct {
	To      string
	Subject string
	Text    string
	HTML    string
}
