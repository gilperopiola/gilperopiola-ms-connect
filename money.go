package connect

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gilperopiola/frutils"
)

type Transaction struct {
	ID          int
	Name        string
	Description string
	Amount      int
	Date        time.Time
}

const msMoneyURL = "http://localhost:9005/v1"

func GetMoneyAmount() (int, error) {
	endpointURL := msMoneyURL + "/Money"

	status, response := frutils.SendHTTPRequest("GET", endpointURL, "")

	if status < 200 || status > 299 {
		return 0, errors.New(response)
	}

	amount := 0
	err := json.Unmarshal([]byte(response), &amount)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func GetTransactions() ([]*Transaction, error) {
	endpointURL := msMoneyURL + "/Transactions"

	status, response := frutils.SendHTTPRequest("GET", endpointURL, "")

	if status < 200 || status > 299 {
		return []*Transaction{}, errors.New(response)
	}

	transactions := []*Transaction{}
	err := json.Unmarshal([]byte(response), &transactions)
	if err != nil {
		return []*Transaction{}, err
	}

	return transactions, nil
}

func GetWeekTransactions() ([]*Transaction, error) {
	endpointURL := msMoneyURL + "/Transactions/Week"

	status, response := frutils.SendHTTPRequest("GET", endpointURL, "")

	if status < 200 || status > 299 {
		return []*Transaction{}, errors.New(response)
	}

	transactions := []*Transaction{}
	err := json.Unmarshal([]byte(response), &transactions)
	if err != nil {
		return []*Transaction{}, err
	}

	return transactions, nil
}

func GetDayTransactions() ([]*Transaction, error) {
	endpointURL := msMoneyURL + "/Transactions/Day"

	status, response := frutils.SendHTTPRequest("GET", endpointURL, "")

	if status < 200 || status > 299 {
		return []*Transaction{}, errors.New(response)
	}

	transactions := []*Transaction{}
	err := json.Unmarshal([]byte(response), &transactions)
	if err != nil {
		return []*Transaction{}, err
	}

	return transactions, nil
}
