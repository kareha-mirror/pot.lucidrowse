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

func updateAreaStateWithOpenAI(
	cfg *config.Config, areaCode string,
) (string, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return "", fmt.Errorf("key not set")
	}

	state, err := data.AreaState(areaCode)
	if err != nil {
		return "", err
	}

	actions, err := data.EventList(areaCode)
	if err != nil {
		return "", err
	}
	actionsStr, err := json.Marshal(actions)
	if err != nil {
		return "", err
	}

	systemMessage :=
		cfg.Prompts.Common + "\n\n" +
			cfg.Prompts.Events + "\n\n" +
			cfg.Prompts.UpdateAreaState
	userMessage :=
		"AREASTATE: " + state + "\n\n" +
			"ACTIONS: " + string(actionsStr)

	// debug
	//fmt.Println("System Message:")
	//fmt.Println(systemMessage)
	//fmt.Println("User Message:")
	//fmt.Println(userMessage)
	fmt.Println("UpdateAreaState: " + areaCode)

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
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}
	text := resp.Choices[0].Message.Content

	return text, nil
}
