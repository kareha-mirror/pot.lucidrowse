package ai

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func updateImageWithOpenAI(cfg *config.Config, current []byte, content string) ([]byte, error) {
	apiKey := cfg.Filter.Key
	if apiKey == "" {
		return []byte{}, fmt.Errorf("Filter key (OpenAI API key) not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	message := "Input: " + content

	resp, err := client.Images.Edit(
		context.Background(),
		openai.ImageEditParams{
			Model:        openai.ImageModelGPTImage1Mini,
			Size:         openai.ImageEditParamsSize1024x1024,
			Quality:      openai.ImageEditParamsQualityLow,
			OutputFormat: openai.ImageEditParamsOutputFormatWebP,
			Prompt:       cfg.Prompts.Common + "\n" + cfg.Prompts.UpdateImage + "\n" + message,
			Image: openai.ImageEditParamsImageUnion{
				OfFile: openai.File(
					bytes.NewReader(current),
					"current.webp",
					"image/webp",
				),
			},
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
