package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func updateWorldStateWithOpenAI(cfg *config.Config) (string, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return "", fmt.Errorf("key not set")
	}

	worldState, err := data.WorldState()
	if err != nil {
		return "", err
	}

	b := strings.Builder{}
	for regionCode, regionName := range data.RegionNames {
		b.WriteString("== " + regionName + " ==\n")
		b.WriteString("\n")

		regionState, err := data.RegionState(regionCode)
		if err != nil {
			return "", err
		}
		b.WriteString(regionState)
		b.WriteString("\n")
	}
	regionStates := b.String()

	systemMessage :=
		cfg.Prompts.Common + "\n\n" +
			cfg.Prompts.UpdateWorldState
	userMessage :=
		"WORLDSTATE: " + worldState + "\n\n" +
			"REGIONSTATES: " + regionStates

	// debug
	//fmt.Println("System Message:")
	//fmt.Println(systemMessage)
	//fmt.Println("User Message:")
	//fmt.Println(userMessage)
	fmt.Println("UpdateWorldState: ")

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
