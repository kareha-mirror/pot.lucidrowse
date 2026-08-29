package ai

import (
	"encoding/json"
	"strings"

	"golang.org/x/text/unicode/norm"
)

type Flavor struct {
	Name        string `json:"name"`
	Race        string `json:"race"`
	Job         string `json:"job"`
	Description string `json:"description"`
	AreaCode    string `json:"area-code"`
	AreaName    string `json:"area-name"`
	Error       string `json:"error"`
}

func sanitizeJSONString(s string) string {
	sanitized := strings.ReplaceAll(s, "```json", "")
	return strings.ReplaceAll(sanitized, "```", "")
}

func ParseFlavor(rawFlavorStr string) (Flavor, error) {
	normFlavorStr := norm.NFC.String(rawFlavorStr)

	flavorStr := sanitizeJSONString(normFlavorStr)
	var flavor Flavor
	if err := json.Unmarshal([]byte(flavorStr), &flavor); err != nil {
		return Flavor{}, err
	}
	return flavor, nil
}
