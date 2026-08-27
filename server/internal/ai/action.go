package ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
	"tea.kareha.org/pot/lucidrowse/server/internal/model"
)

func actionWithOpenAI(
	cfg *config.Config, flavor model.Flavor, input string,
) (model.Action, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return model.Action{}, fmt.Errorf("key not set")
	}

	flavorStr, err := flavor.Marshal()
	if err != nil {
		return model.Action{}, err
	}

	actions, err := data.EventList(flavor.AreaCode)
	if err != nil {
		return model.Action{}, err
	}
	actionsStr, err := json.Marshal(actions)
	if err != nil {
		return model.Action{}, err
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	systemMessage := cfg.Prompts.Common + "\n" + cfg.Prompts.Events + "\n" + cfg.Prompts.Action
	userMessage := "ACTIONS: " + string(actionsStr) + "\nJSON: " + flavorStr + "\nInput: " + input

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
		return model.Action{}, err
	}

	if len(resp.Choices) == 0 {
		return model.Action{}, fmt.Errorf("no choices in response")
	}
	text := resp.Choices[0].Message.Content

	return model.ParseAction(text)
}
