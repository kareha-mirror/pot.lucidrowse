package filter

import (
	"fmt"

	"golang.org/x/text/unicode/norm"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func Apply(cfg *config.Config, content string) (string, error) {
	normContent := norm.NFC.String(content)

	if cfg.Filter.Agent == "openai" {
		filtered, err := withOpenAI(cfg, normContent)
		return norm.NFC.String(filtered), err
	} else if cfg.Filter.Agent == "nil" {
		return normContent, nil
	} else {
		return "", fmt.Errorf("Invalid filter agent. If you want to disable filter, set it to \"nil\".")
	}
}

func ApplyImage(cfg *config.Config, content string) ([]byte, error) {
	normContent := norm.NFC.String(content)

	if cfg.Filter.Agent == "openai" {
		return ImageWithOpenAI(cfg, normContent)
	} else if cfg.Filter.Agent == "nil" {
		return []byte{}, nil
	} else {
		return []byte{}, fmt.Errorf("Invalid filter agent. If you want to disable filter, set it to \"nil\".")
	}
}
