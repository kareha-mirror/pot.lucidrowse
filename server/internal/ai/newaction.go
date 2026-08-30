package ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func newActionWithOpenAI(
	cfg *config.Config, flavor Flavor, input string,
) (Action, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return Action{}, fmt.Errorf("key not set")
	}

	flavorStr, err := json.Marshal(flavor)
	if err != nil {
		return Action{}, err
	}

	areaState, err := data.AreaState(flavor.AreaCode)
	if err != nil {
		return Action{}, err
	}

	regionCode, err := data.RegionFromArea(flavor.AreaCode)
	if err != nil {
		return Action{}, err
	}
	regionState, err := data.RegionState(regionCode)
	if err != nil {
		return Action{}, err
	}

	worldState, err := data.WorldState()
	if err != nil {
		return Action{}, err
	}

	actions, err := data.EventList(flavor.AreaCode)
	if err != nil {
		return Action{}, err
	}
	actionsStr, err := json.Marshal(actions)
	if err != nil {
		return Action{}, err
	}

	systemMessage :=
		cfg.Prompts.Common + "\n\n" +
			cfg.Prompts.Events + "\n\n" +
			cfg.Prompts.Action
	userMessage :=
		"AREASTATE: " + areaState + "\n\n" +
			"REGIONSTATE: " + regionState + "\n\n" +
			"WORLDSTATE: " + worldState + "\n\n" +
			"ACTIONS: " + string(actionsStr) + "\n\n" +
			"FLAVOR: " + string(flavorStr) + "\n\n" +
			"INPUT: " + input

	// debug
	//fmt.Println("System Message:")
	//fmt.Println(systemMessage)
	//fmt.Println("User Message:")
	//fmt.Println(userMessage)

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	resp, err := client.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Model: openai.ChatModelGPT4o,
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(systemMessage),
				openai.UserMessage(userMessage),
			},
			Temperature: openai.Float(cfg.AI.Temperature),
			TopP:        openai.Float(cfg.AI.TopP),
		},
	)
	if err != nil {
		return Action{}, err
	}

	if len(resp.Choices) == 0 {
		return Action{}, fmt.Errorf("no choices in response")
	}
	actionStr := resp.Choices[0].Message.Content

	return ParseAction(actionStr)
}
