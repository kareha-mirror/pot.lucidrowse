package filter

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func ImageWithOpenAI(cfg *config.Config, content string) ([]byte, error) {
	apiKey := cfg.Filter.Key
	if apiKey == "" {
		return []byte{}, fmt.Errorf("Filter key (OpenAI API key) not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	message := "Input: " + content

	resp, err := client.Images.Generate(
		context.Background(),
		openai.ImageGenerateParams{
			Model:        openai.ImageModelGPTImage1,
			OutputFormat: openai.ImageGenerateParamsOutputFormatWebP,
			Prompt:       cfg.Prompts.Common + "\n" + cfg.Prompts.Image + "\n" + message,
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
