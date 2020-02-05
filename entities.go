package connect

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gilperopiola/frutils"
)

type Entity struct {
	ID          int
	Name        string
	Description string
	Kind        string
	Importance  int
	Status      int //1 = enabled, 2 = disabled
	DateCreated time.Time
}

const msEntitiesURL = "http://localhost:9001"

func GetEntitiesOfKind(kind string) ([]*Entity, error) {
	endpointURL := msEntitiesURL + "/Entity?kind=" + kind

	status, response := frutils.SendHTTPRequest("GET", endpointURL, "")

	if status < 200 || status > 299 {
		return []*Entity{}, errors.New(response)
	}

	entities := []*Entity{}
	err := json.Unmarshal([]byte(response), &entities)
	if err != nil {
		return []*Entity{}, err
	}

	return entities, nil
}
