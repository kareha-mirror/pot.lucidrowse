package ai

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func newFlavorImageWithOpenAI(
	cfg *config.Config, flavor Flavor,
) (data.Image, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return data.Image{}, fmt.Errorf("key not set")
	}

	flavorStr, err := json.Marshal(flavor)
	if err != nil {
		return data.Image{}, err
	}

	message :=
		"FLAVOR: " + string(flavorStr)
	prompt :=
		cfg.Prompts.Common + "\n\n" +
			cfg.Prompts.Image + "\n\n" +
			message

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

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
		return data.Image{}, err
	}

	if len(resp.Data) == 0 {
		return data.Image{}, fmt.Errorf("no choices in response")
	}
	image, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		return data.Image{}, err
	}

	return data.Image{ContentType: "image/webp", Content: image}, nil
}
