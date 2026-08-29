package ai

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func updateFlavorImageWithOpenAI(
	cfg *config.Config, image data.Image, updatedFlavor Flavor,
) (data.Image, error) {
	apiKey := cfg.AI.Key
	if apiKey == "" {
		return data.Image{}, fmt.Errorf("key not set")
	}

	updatedFlavorStr, err := json.Marshal(updatedFlavor)
	if err != nil {
		return data.Image{}, err
	}

	message :=
		"FLAVOR: " + string(updatedFlavorStr)
	prompt :=
		cfg.Prompts.Common + "\n\n" +
			cfg.Prompts.UpdateImage + "\n\n" +
			message

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

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
	updatedImage, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		return data.Image{}, nil
	}

	return data.Image{ContentType: "image/webp", Content: updatedImage}, nil
}
