package ai

import (
	"fmt"

	"golang.org/x/text/unicode/norm"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func NewFlavor(cfg *config.Config, input string) (Flavor, error) {
	normInput := norm.NFC.String(input)

	switch cfg.AI.Agent {
	case "openai":
		return newFlavorWithOpenAI(cfg, normInput)
	default:
		return Flavor{}, fmt.Errorf("invalid AI agent")
	}
}

func NewFlavorImage(cfg *config.Config, flavor Flavor) (data.Image, error) {
	switch cfg.AI.Agent {
	case "openai":
		return newFlavorImageWithOpenAI(cfg, flavor)
	default:
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func UpdateFlavor(
	cfg *config.Config, flavor Flavor, input string,
) (Flavor, error) {
	normInput := norm.NFC.String(input)

	switch cfg.AI.Agent {
	case "openai":
		return updateFlavorWithOpenAI(cfg, flavor, normInput)
	default:
		return Flavor{}, fmt.Errorf("invalid AI agent")
	}
}

func UpdateFlavorImage(
	cfg *config.Config, image data.Image, newFlavor Flavor,
) (data.Image, error) {
	switch cfg.AI.Agent {
	case "openai":
		return updateFlavorImageWithOpenAI(cfg, image, newFlavor)
	default:
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func NewAction(
	cfg *config.Config, flavor Flavor, input string,
) (Action, error) {
	normInput := norm.NFC.String(input)

	switch cfg.AI.Agent {
	case "openai":
		return newActionWithOpenAI(cfg, flavor, normInput)
	default:
		return Action{}, fmt.Errorf("invalid AI agent")
	}
}

func NewActionImage(
	cfg *config.Config, image data.Image, action Action,
) (data.Image, error) {
	switch cfg.AI.Agent {
	case "openai":
		return newActionImageWithOpenAI(cfg, image, action)
	default:
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func UpdateAreaState(cfg *config.Config, areaCode string) (string, error) {
	switch cfg.AI.Agent {
	case "openai":
		return updateAreaStateWithOpenAI(cfg, areaCode)
	default:
		return "", fmt.Errorf("invalid AI agent")
	}
}

func UpdateRegionState(cfg *config.Config, regionCode string) (string, error) {
	switch cfg.AI.Agent {
	case "openai":
		return updateRegionStateWithOpenAI(cfg, regionCode)
	default:
		return "", fmt.Errorf("invalid AI agent")
	}
}

func UpdateWorldState(cfg *config.Config) (string, error) {
	switch cfg.AI.Agent {
	case "openai":
		return updateWorldStateWithOpenAI(cfg)
	default:
		return "", fmt.Errorf("invalid AI agent")
	}
}
