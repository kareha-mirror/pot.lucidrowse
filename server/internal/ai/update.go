package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
	"tea.kareha.org/pot/lucidrowse/server/internal/model"
)

func updateWithOpenAI(
	cfg *config.Config, flavor model.Flavor, input string,
) (model.Flavor, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return model.Flavor{}, fmt.Errorf("key not set")
	}

	flavorStr, err := flavor.Marshal()
	if err != nil {
		return model.Flavor{}, err
	}

	areas, err := data.DescribeAreas()
	if err != nil {
		return model.Flavor{}, err
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	systemMessage := cfg.Prompts.Common + "\n" + cfg.Prompts.Areas + areas + "\n" + cfg.Prompts.Update
	userMessage := "JSON: " + flavorStr + "\nInput: " + input

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
		return model.Flavor{}, err
	}

	if len(resp.Choices) == 0 {
		return model.Flavor{}, fmt.Errorf("no choices in response")
	}
	updatedFlavor := resp.Choices[0].Message.Content

	return model.ParseFlavor(updatedFlavor)
}
