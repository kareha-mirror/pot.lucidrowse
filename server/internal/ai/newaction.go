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

	flavorStr, err := flavor.Marshal()
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

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	systemMessage := cfg.Prompts.Common + "\n" + cfg.Prompts.Events + "\n" + cfg.Prompts.Action
	userMessage := "ACTIONS: " + string(actionsStr) + "\nJSON: " + flavorStr + "\nInput: " + input

	// debug
	//fmt.Println("System Message:")
	//fmt.Println(systemMessage)
	//fmt.Println("User Message:")
	//fmt.Println(userMessage)

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
	text := resp.Choices[0].Message.Content

	return ParseAction(text)
}
