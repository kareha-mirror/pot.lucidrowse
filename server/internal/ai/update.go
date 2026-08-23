package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func updateWithOpenAI(
	cfg *config.Config, flavor string, input string,
) (string, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return "", fmt.Errorf("key not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	systemMessage := cfg.Prompts.Common + "\n" + cfg.Prompts.Update
	userMessage := "JSON: " + flavor + "\nInput: " + input

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
	updatedFlavor := resp.Choices[0].Message.Content

	return updatedFlavor, nil
}
