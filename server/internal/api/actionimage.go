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

func actionImage(cfg *config.Config, current []byte, text string) ([]byte, error) {
	return ai.ActionImage(cfg, current, text)
}

func handleActionImage(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	var req ImageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	text, ok := actions[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	current, ok := images[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	image, err := actionImage(cfg, current, text)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/webp")
	w.Write(image)
}

func wrapActionImageHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleActionImage(cfg, w, r)
	}
}
