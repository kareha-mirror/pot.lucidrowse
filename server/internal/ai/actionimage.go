package ai

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func actionImageWithOpenAI(
	cfg *config.Config, image data.Image, text string,
) (data.Image, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return data.Image{}, fmt.Errorf("key not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	message := "Text: " + text
	prompt := cfg.Prompts.Common + "\n" + cfg.Prompts.ActionImage + "\n" + message

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
					bytes.NewReader(image.Content),
					"image.webp",
					image.ContentType,
				),
			},
		},
	)
	if err != nil {
		return data.Image{}, err
	}

	if len(resp.Data) == 0 {
		return data.Image{}, fmt.Errorf("no choices in response")
	}
	newImage, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)

	return data.Image{ContentType: "image/webp", Content: newImage}, nil
}
