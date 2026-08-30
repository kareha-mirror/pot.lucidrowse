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

func updateRegionStateWithOpenAI(
	cfg *config.Config, regionCode string,
) (string, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return "", fmt.Errorf("key not set")
	}

	regionState, err := data.RegionState(regionCode)
	if err != nil {
		return "", err
	}

	areas, err := data.AreaList()
	if err != nil {
		return "", err
	}

	b := strings.Builder{}
	for _, area := range areas {
		if area.RegionCode != regionCode {
			continue
		}

		b.WriteString("== " + data.RegionNames[area.RegionCode] + " " + area.Name + " ==\n")
		b.WriteString("\n")

		areaState, err := data.AreaState(area.RegionCode + "-" + area.AreaCode)
		if err != nil {
			return "", err
		}
		b.WriteString(areaState)
		b.WriteString("\n")
	}
	areaStates := b.String()

	systemMessage :=
		cfg.Prompts.Common + "\n\n" +
			cfg.Prompts.UpdateRegionState
	userMessage :=
		"REGIONSTATE: " + regionState + "\n\n" +
			"AREASTATES: " + areaStates

	// debug
	//fmt.Println("System Message:")
	//fmt.Println(systemMessage)
	//fmt.Println("User Message:")
	//fmt.Println(userMessage)
	fmt.Println("UpdateRegionState: " + regionCode)

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
