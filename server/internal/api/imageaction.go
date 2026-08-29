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

type ImageActionResponse struct {
	ImageId string `json:"image-id"`
}

func handleImageAction(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	playerPubId := r.PathValue("id")

	playerId, err := data.PlayerId(playerPubId)
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

	f, err := data.LoadCurrentFlavor(playerId)
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

	action := ai.Action{Description: a.Description}
	newImage, err := ai.NewActionImage(cfg, image, action)
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
