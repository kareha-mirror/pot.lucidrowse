package ai

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func imageWithOpenAI(cfg *config.Config, flavor string) ([]byte, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return []byte{}, fmt.Errorf("key not set")
	}

	areas, err := data.DescribeAreas()
	if err != nil {
		return []byte{}, err
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	message := "JSON: " + flavor
	prompt := cfg.Prompts.Common + "\n" + areas + "\n" + cfg.Prompts.Image + "\n" + message

	resp, err := client.Images.Generate(
		context.Background(),
		openai.ImageGenerateParams{
			Model:        openai.ImageModelGPTImage1Mini,
			Size:         openai.ImageGenerateParamsSize1024x1024,
			Quality:      openai.ImageGenerateParamsQualityLow,
			OutputFormat: openai.ImageGenerateParamsOutputFormatWebP,
			Prompt:       prompt,
		},
	)
	if err != nil {
		return []byte{}, err
	}

	if len(resp.Data) == 0 {
		return []byte{}, fmt.Errorf("no choices in response")
	}
	image, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)

	return image, nil
}
