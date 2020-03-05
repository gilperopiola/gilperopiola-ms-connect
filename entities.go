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

func CreateEntity(name string, description string, kind string, importance int) (*Entity, error) {
	endpointURL := msEntitiesURL + "/Entity"

	httpRequestBody := `{
		"name": "` + name + `",
		"description": "` + description + `",
		"kind": "` + kind + `",
		"importance": ` + frutils.ToString(importance) + `
	}`

	status, response := frutils.SendHTTPRequest("POST", endpointURL, httpRequestBody)

	if status < 200 || status > 299 {
		return &Entity{}, errors.New(response)
	}

	entity := &Entity{}
	err := json.Unmarshal([]byte(response), &entity)
	if err != nil {
		return &Entity{}, err
	}

	return entity, nil
}

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
