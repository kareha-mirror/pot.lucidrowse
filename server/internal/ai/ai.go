package ai

import (
	"fmt"

	"golang.org/x/text/unicode/norm"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func Create(cfg *config.Config, input string) (string, error) {
	normContent := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		flavor, err := createWithOpenAI(cfg, normContent)
		return norm.NFC.String(flavor), err
	} else {
		return "", fmt.Errorf("invalid AI agent")
	}
}

func Image(cfg *config.Config, flavor string) (data.Image, error) {
	normContent := norm.NFC.String(flavor)

	if cfg.AI.Agent == "openai" {
		return imageWithOpenAI(cfg, normContent)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func Update(cfg *config.Config, flavor string, input string) (string, error) {
	normCurrent := norm.NFC.String(flavor)
	normContent := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		updatedFlavor, err := updateWithOpenAI(cfg, normCurrent, normContent)
		return norm.NFC.String(updatedFlavor), err
	} else {
		return "", fmt.Errorf("invalid AI agent")
	}
}

func UpdateImage(
	cfg *config.Config, image data.Image, newFlavor string,
) (data.Image, error) {
	normContent := norm.NFC.String(newFlavor)

	if cfg.AI.Agent == "openai" {
		return updateImageWithOpenAI(cfg, image, normContent)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func Action(cfg *config.Config, flavor string, input string) (string, error) {
	normCurrent := norm.NFC.String(flavor)
	normContent := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		text, err := actionWithOpenAI(cfg, normCurrent, normContent)
		return norm.NFC.String(text), err
	} else {
		return "", fmt.Errorf("invalid AI agent")
	}
}

func ActionImage(
	cfg *config.Config, image data.Image, text string,
) (data.Image, error) {
	normContent := norm.NFC.String(text)

	if cfg.AI.Agent == "openai" {
		return actionImageWithOpenAI(cfg, image, normContent)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}
