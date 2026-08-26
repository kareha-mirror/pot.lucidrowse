package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ImageActionRequest struct {
	PlayerId string `json:"player-id"`
}

type ImageActionResponse struct {
	ImageId string `json:"image-id"`
}

func newActionImage(
	cfg *config.Config, image data.Image, text string,
) (data.Image, error) {
	return ai.ActionImage(cfg, image, text)
}

func handleImageAction(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	var req ImageActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	playerId, err := data.PlayerId(req.PlayerId)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	a, err := data.LoadLastAction(playerId)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	f, err := data.LoadLastFlavor(playerId)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if f.ImagePubId == nil {
		log.Println("image not found")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	image, err := data.LoadImage(context.Background(), *f.ImagePubId)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	newImage, err := newActionImage(cfg, image, a.Description)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	imagePubId, err := newPubId()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = data.SaveImage(context.Background(), imagePubId, newImage)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = data.AddImageToLastAction(playerId, imagePubId)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := ImageActionResponse{
		ImageId: imagePubId,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func imageActionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleImageAction(cfg, w, r)
	}
}
