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

type ImageFlavorResponse struct {
	ImageId string `json:"image-id"`
}

func handleImageFlavor(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	playerPubId := r.PathValue("id")

	playerId, err := data.PlayerId(playerPubId)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	f, err := data.LoadLastFlavor(playerId)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	flavor := ai.Flavor{
		Name:        f.Name,
		Race:        f.Race,
		Job:         f.Job,
		Description: f.Description,
		AreaCode:    f.AreaCode,
		AreaName:    f.AreaName,
	}

	var newImage data.Image
	if f.ImagePubId == nil {
		newImage, err = ai.NewFlavorImage(cfg, flavor)
	} else {
		image, err := data.LoadImage(context.Background(), *f.ImagePubId)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		newImage, err = ai.UpdateFlavorImage(cfg, image, flavor)
	}
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

	err = data.AddImageToLastFlavor(playerId, imagePubId)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := ImageFlavorResponse{
		ImageId: imagePubId,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func imageFlavorHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleImageFlavor(cfg, w, r)
	}
}
