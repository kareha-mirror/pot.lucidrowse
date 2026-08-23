package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func actionWithOpenAI(cfg *config.Config, current string, content string) (string, error) {
	apiKey := cfg.Filter.Key
	if apiKey == "" {
		return "", fmt.Errorf("Filter key (OpenAI API key) not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	message := "JSON: " + current + "\nInput: " + content

	resp, err := client.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Model: openai.ChatModelGPT4o,
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(cfg.Prompts.Common + "\n" + cfg.Prompts.Action),
				openai.UserMessage(message),
			},
			Temperature: openai.Float(cfg.Filter.Temperature),
			TopP:        openai.Float(cfg.Filter.TopP),
		},
	)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}
	answer := resp.Choices[0].Message.Content

	return answer, nil
}
