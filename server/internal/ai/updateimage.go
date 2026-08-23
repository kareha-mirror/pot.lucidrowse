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

func updateImageWithOpenAI(
	cfg *config.Config, image []byte, updatedFlavor string,
) ([]byte, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return []byte{}, fmt.Errorf("key not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	message := "Input: " + updatedFlavor
	prompt := cfg.Prompts.Common + "\n" + cfg.Prompts.UpdateImage + "\n" + message

	resp, err := client.Images.Edit(
		context.Background(),
		openai.ImageEditParams{
			Model:        openai.ImageModelGPTImage1Mini,
			Size:         openai.ImageEditParamsSize1024x1024,
			Quality:      openai.ImageEditParamsQualityLow,
			OutputFormat: openai.ImageEditParamsOutputFormatWebP,
			Prompt:       prompt,
			Image: openai.ImageEditParamsImageUnion{
				OfFile: openai.File(
					bytes.NewReader(image),
					"image.webp",
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
	updatedImage, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)

	return updatedImage, nil
}
