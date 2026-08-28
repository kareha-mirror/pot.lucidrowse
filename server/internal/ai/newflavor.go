package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func newFlavorWithOpenAI(cfg *config.Config, input string) (Flavor, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return Flavor{}, fmt.Errorf("key not set")
	}

	areas, err := data.DescribeAreas()
	if err != nil {
		return Flavor{}, err
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	systemMessage := cfg.Prompts.Common + "\n" + cfg.Prompts.Areas + areas + "\n" + cfg.Prompts.Create
	userMessage := "Input: " + input

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
		return Flavor{}, err
	}

	if len(resp.Choices) == 0 {
		return Flavor{}, fmt.Errorf("no choices in response")
	}
	flavor := resp.Choices[0].Message.Content

	return ParseFlavor(flavor)
}
