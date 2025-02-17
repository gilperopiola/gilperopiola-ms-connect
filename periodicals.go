package connect

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gilperopiola/frutils"
)

type Periodical struct {
	ID         int
	Name       string
	Days       int
	LastDone   time.Time
	Importance int
	Color      string
	Enabled    bool
	Archived   bool
}

const msPeriodicalsURL = "http://localhost:9002"

func CreatePeriodical(name string, days int, importance int, color string) (*Periodical, error) {
	endpointURL := msPeriodicalsURL + "/Periodical"

	httpRequestBody := `{
		"name": "` + name + `",
		"days": ` + frutils.ToString(days) + `,
		"importance": ` + frutils.ToString(importance) + `,
		"color": "` + color + `"
	}`

	status, response := frutils.SendHTTPRequest("POST", endpointURL, httpRequestBody)

	if status < 200 || status > 299 {
		return &Periodical{}, errors.New(response)
	}

	periodical := &Periodical{}
	json.Unmarshal([]byte(response), periodical)

	return periodical, nil
}

func GetPeriodical(id int) (*Periodical, error) {
	endpointURL := msPeriodicalsURL + "/Periodical/" + frutils.ToString(id)

	status, response := frutils.SendHTTPRequest("GET", endpointURL, "")

	if status < 200 || status > 299 {
		return &Periodical{}, errors.New(response)
	}

	periodical := &Periodical{}
	json.Unmarshal([]byte(response), periodical)

	return periodical, nil
}

func UpdatePeriodical(id int, name string, days int, importance int, color string, enabled bool, lastDone time.Time) (*Periodical, error) {
	endpointURL := msPeriodicalsURL + "/Periodical/" + frutils.ToString(id)

	httpRequestBody := `{
		"name": "` + name + `",
		"days": ` + frutils.ToString(days) + `,
		"importance": ` + frutils.ToString(importance) + `,
		"color": "` + color + `",
		"enabled": ` + frutils.BoolToString(enabled) + `,
		"lastDone": "` + lastDone.Format(time.RFC3339) + `"
	}`

	status, response := frutils.SendHTTPRequest("PUT", endpointURL, httpRequestBody)

	if status < 200 || status > 299 {
		return &Periodical{}, errors.New(response)
	}

	periodical := &Periodical{}
	json.Unmarshal([]byte(response), periodical)

	return periodical, nil
}

func GetAllPeriodicals() ([]*Periodical, error) {
	endpointURL := msPeriodicalsURL + "/Periodical?filterEnabled=true&filterArchived=false"

	status, response := frutils.SendHTTPRequest("GET", endpointURL, "")

	if status < 200 || status > 299 {
		return []*Periodical{}, errors.New(response)
	}

	periodicals := []*Periodical{}
	json.Unmarshal([]byte(response), &periodicals)

	return periodicals, nil
}

/* mailing */

func GetPeriodicalsExpiringToday() ([]*Periodical, error) {
	periodicals, err := GetAllPeriodicals()
	if err != nil {
		return []*Periodical{}, err
	}

	filteredPeriodicals := []*Periodical{}
	for _, periodical := range periodicals {
		daysDiff := frutils.GetDaysBetween(time.Now(), periodical.LastDone)

		if periodical.Days-daysDiff <= 0 {
			filteredPeriodicals = append(filteredPeriodicals, periodical)
		}
	}

	return filteredPeriodicals, nil
}

func GetPeriodicalsExpiringTomorrow() ([]*Periodical, error) {
	periodicals, err := GetAllPeriodicals()
	if err != nil {
		return []*Periodical{}, err
	}

	filteredPeriodicals := []*Periodical{}
	for _, periodical := range periodicals {
		daysDiff := frutils.GetDaysBetween(time.Now(), periodical.LastDone)

		if periodical.Days-daysDiff == 1 {
			filteredPeriodicals = append(filteredPeriodicals, periodical)
		}
	}

	return filteredPeriodicals, nil
}

func GetPeriodicalsDoneYesterday() ([]*Periodical, error) {
	periodicals, err := GetAllPeriodicals()
	if err != nil {
		return []*Periodical{}, err
	}

	filteredPeriodicals := []*Periodical{}
	for _, periodical := range periodicals {
		daysDiff := frutils.GetDaysBetween(time.Now(), periodical.LastDone)

		if daysDiff <= 1 {
			filteredPeriodicals = append(filteredPeriodicals, periodical)
		}
	}

	return filteredPeriodicals, nil
}
