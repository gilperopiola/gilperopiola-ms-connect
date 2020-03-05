package connect

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/gilperopiola/frutils"
)

type Task struct {
	ID          int
	Name        string
	Description string
	Importance  int
	Duration    TaskDuration
	DueDate     time.Time
}

type TaskDuration int

const (
	ExtraSmall TaskDuration = 1
	Small      TaskDuration = 2
	Medium     TaskDuration = 3
	Large      TaskDuration = 4
	ExtraLarge TaskDuration = 5
)

type User struct {
	Token string
}

const lyfeCompanyonURL = "http://localhost:9001"

func GetLyfeCompanyonToken(username string, password string) string {
	endpointURL := lyfeCompanyonURL + "/Login"

	httpRequestBody := `{
		"username": "` + username + `",
		"password": "` + password + `"
	}`

	status, response := frutils.SendHTTPRequest("POST", endpointURL, httpRequestBody)

	if status < 200 || status > 299 {
		return ""
	}

	user := &User{}
	json.Unmarshal([]byte(response), user)

	return user.Token
}

func CreateTask(name string, importance int, duration int, daily bool, weekly bool, monthly bool, token string) (*Task, error) {
	endpointURL := lyfeCompanyonURL + "/Tasks"

	tagID := 0
	if daily {
		tagID = 1
	}
	if weekly {
		tagID = 2
	}
	if monthly {
		tagID = 4
	}

	log.Println(frutils.ToString(tagID))

	httpRequestBody := `{
		"name": "` + name + `",
		"duration": ` + frutils.ToString(duration) + `,
		"importance": ` + frutils.ToString(importance) + `
	}`

	status, response := frutils.SendHTTPRequestWithToken("POST", endpointURL, httpRequestBody, token)

	if status < 200 || status > 299 {
		return &Task{}, errors.New(response)
	}

	task := &Task{}
	json.Unmarshal([]byte(response), task)

	return task, nil
}
