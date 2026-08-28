package ai

import (
	"fmt"

	"golang.org/x/text/unicode/norm"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func NewFlavor(cfg *config.Config, input string) (Flavor, error) {
	normInput := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		return newFlavorWithOpenAI(cfg, normInput)
	} else {
		return Flavor{}, fmt.Errorf("invalid AI agent")
	}
}

func NewFlavorImage(cfg *config.Config, flavor Flavor) (data.Image, error) {
	if cfg.AI.Agent == "openai" {
		return newFlavorImageWithOpenAI(cfg, flavor)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func UpdateFlavor(
	cfg *config.Config, flavor Flavor, input string,
) (Flavor, error) {
	normInput := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		return updateFlavorWithOpenAI(cfg, flavor, normInput)
	} else {
		return Flavor{}, fmt.Errorf("invalid AI agent")
	}
}

func UpdateFlavorImage(
	cfg *config.Config, image data.Image, newFlavor Flavor,
) (data.Image, error) {
	if cfg.AI.Agent == "openai" {
		return updateFlavorImageWithOpenAI(cfg, image, newFlavor)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func NewAction(
	cfg *config.Config, flavor Flavor, input string,
) (Action, error) {
	normInput := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		return newActionWithOpenAI(cfg, flavor, normInput)
	} else {
		return Action{}, fmt.Errorf("invalid AI agent")
	}
}

func NewActionImage(
	cfg *config.Config, image data.Image, action Action,
) (data.Image, error) {
	if cfg.AI.Agent == "openai" {
		return newActionImageWithOpenAI(cfg, image, action)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func UpdateState(cfg *config.Config, areaCode string) (string, error) {
	if cfg.AI.Agent == "openai" {
		return updateStateWithOpenAI(cfg, areaCode)
	} else {
		return "", fmt.Errorf("invalid AI agent")
	}
}
