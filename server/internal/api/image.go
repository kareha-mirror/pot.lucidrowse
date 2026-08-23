package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

type ImageRequest struct {
	Id string `json:"id"`
}

func generateImage(cfg *config.Config, flavor string) ([]byte, error) {
	return ai.Image(cfg, flavor)
}

func updateImage(
	cfg *config.Config, image []byte, updatedFlavor string,
) ([]byte, error) {
	return ai.UpdateImage(cfg, image, updatedFlavor)
}

var images = map[string][]byte{}

func handleImage(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	var req ImageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	flavor, ok := flavors[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	image, ok := images[req.Id]
	var newImage []byte
	var err error
	if ok {
		newImage, err = updateImage(cfg, image, flavor)
	} else {
		newImage, err = generateImage(cfg, flavor)
	}
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	images[req.Id] = newImage

	w.Header().Set("Content-Type", "image/webp")
	w.Write(newImage)
}

func wrapImageHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleImage(cfg, w, r)
	}
}
