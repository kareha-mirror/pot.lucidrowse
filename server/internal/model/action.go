package model

import (
	"encoding/json"

	"golang.org/x/text/unicode/norm"
)

type Action struct {
	Description string `json:"description"`
	Error       string `json:"error"`
}

func ParseAction(rawActionStr string) (Action, error) {
	normRawActionStr := norm.NFC.String(rawActionStr)

	actionStr := sanitizeJSONString(normRawActionStr)
	var action Action
	if err := json.Unmarshal([]byte(actionStr), &action); err != nil {
		return Action{}, err
	}
	return action, nil
}
