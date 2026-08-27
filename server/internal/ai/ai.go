package ai

import (
	"fmt"

	"golang.org/x/text/unicode/norm"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
	"tea.kareha.org/pot/lucidrowse/server/internal/model"
)

func Create(cfg *config.Config, input string) (model.Flavor, error) {
	normInput := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		return createWithOpenAI(cfg, normInput)
	} else {
		return model.Flavor{}, fmt.Errorf("invalid AI agent")
	}
}

func Image(cfg *config.Config, flavor model.Flavor) (data.Image, error) {
	if cfg.AI.Agent == "openai" {
		return imageWithOpenAI(cfg, flavor)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func Update(
	cfg *config.Config, flavor model.Flavor, input string,
) (model.Flavor, error) {
	normInput := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		return updateWithOpenAI(cfg, flavor, normInput)
	} else {
		return model.Flavor{}, fmt.Errorf("invalid AI agent")
	}
}

func UpdateImage(
	cfg *config.Config, image data.Image, newFlavor model.Flavor,
) (data.Image, error) {
	if cfg.AI.Agent == "openai" {
		return updateImageWithOpenAI(cfg, image, newFlavor)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}

func Action(
	cfg *config.Config, flavor model.Flavor, input string,
) (model.Action, error) {
	normInput := norm.NFC.String(input)

	if cfg.AI.Agent == "openai" {
		return actionWithOpenAI(cfg, flavor, normInput)
	} else {
		return model.Action{}, fmt.Errorf("invalid AI agent")
	}
}

func ActionImage(
	cfg *config.Config, image data.Image, action model.Action,
) (data.Image, error) {
	if cfg.AI.Agent == "openai" {
		return actionImageWithOpenAI(cfg, image, action)
	} else {
		return data.Image{}, fmt.Errorf("invalid AI agent")
	}
}
