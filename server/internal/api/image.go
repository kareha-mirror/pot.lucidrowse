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

func generateImage(cfg *config.Config, text string) ([]byte, error) {
	return ai.Image(cfg, text)
}

func updateImage(cfg *config.Config, current []byte, text string) ([]byte, error) {
	return ai.UpdateImage(cfg, current, text)
}

var images = map[string][]byte{}

func handleImage(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	var req ImageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	text, ok := texts[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	current, ok := images[req.Id]
	var image []byte
	var err error
	if ok {
		image, err = updateImage(cfg, current, text)
	} else {
		image, err = generateImage(cfg, text)
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	images[req.Id] = image

	w.Header().Set("Content-Type", "image/webp")
	w.Write(image)
}

func wrapImageHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleImage(cfg, w, r)
	}
}
