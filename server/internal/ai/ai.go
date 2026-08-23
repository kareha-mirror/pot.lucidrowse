package ai

import (
	"fmt"

	"golang.org/x/text/unicode/norm"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func Create(cfg *config.Config, content string) (string, error) {
	normContent := norm.NFC.String(content)

	if cfg.Filter.Agent == "openai" {
		filtered, err := createWithOpenAI(cfg, normContent)
		return norm.NFC.String(filtered), err
	} else {
		return "", fmt.Errorf("invalid AI agent")
	}
}

func Image(cfg *config.Config, content string) ([]byte, error) {
	normContent := norm.NFC.String(content)

	if cfg.Filter.Agent == "openai" {
		return imageWithOpenAI(cfg, normContent)
	} else {
		return []byte{}, fmt.Errorf("invalid AI agent")
	}
}

func Update(cfg *config.Config, current string, content string) (string, error) {
	normCurrent := norm.NFC.String(current)
	normContent := norm.NFC.String(content)

	if cfg.Filter.Agent == "openai" {
		filtered, err := updateWithOpenAI(cfg, normCurrent, normContent)
		return norm.NFC.String(filtered), err
	} else {
		return "", fmt.Errorf("invalid AI agent")
	}
}

func UpdateImage(cfg *config.Config, current []byte, content string) ([]byte, error) {
	normContent := norm.NFC.String(content)

	if cfg.Filter.Agent == "openai" {
		return updateImageWithOpenAI(cfg, current, normContent)
	} else {
		return []byte{}, fmt.Errorf("invalid AI agent")
	}
}

func Action(cfg *config.Config, current string, content string) (string, error) {
	normCurrent := norm.NFC.String(current)
	normContent := norm.NFC.String(content)

	if cfg.Filter.Agent == "openai" {
		filtered, err := actionWithOpenAI(cfg, normCurrent, normContent)
		return norm.NFC.String(filtered), err
	} else {
		return "", fmt.Errorf("invalid AI agent")
	}
}

func ActionImage(cfg *config.Config, current []byte, content string) ([]byte, error) {
	normContent := norm.NFC.String(content)

	if cfg.Filter.Agent == "openai" {
		return actionImageWithOpenAI(cfg, current, normContent)
	} else {
		return []byte{}, fmt.Errorf("invalid AI agent")
	}
}
