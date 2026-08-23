package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

type ActionImageRequest struct {
	Id string `json:"id"`
}

func actionImage(
	cfg *config.Config, image []byte, text string,
) ([]byte, error) {
	return ai.ActionImage(cfg, image, text)
}

func handleActionImage(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	var req ActionImageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	text, ok := actions[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	image, ok := images[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	newImage, err := actionImage(cfg, image, text)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/webp")
	w.Write(newImage)
}

func wrapActionImageHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleActionImage(cfg, w, r)
	}
}
