package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ImageActionResponse struct {
	ImageID string `json:"image-id"`
}

func (api *API) handleImageAction(w http.ResponseWriter, r *http.Request) {
	user, player, err := authPlayer(w, r)
	if err != nil {
		return
	}

	if user.AICalls >= api.cfg.Game.MaxAICalls {
		http.Error(w, "too many requests", http.StatusTooManyRequests)
		return
	}
	if err = data.IncrementAICalls(user.ID); err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to increment AI calls",
			http.StatusInternalServerError,
		)
		return
	}

	a, err := data.LoadLastAction(player.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "last action not found", http.StatusNotFound)
		return
	}

	f, err := data.LoadCurrentFlavor(player.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "current flavor not found", http.StatusNotFound)
		return
	}

	if f.ImagePubID == nil {
		log.Println("image not found")
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}
	image, err := data.LoadImage(context.Background(), *f.ImagePubID)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to load image", http.StatusInternalServerError)
		return
	}

	action := ai.Action{Description: a.Description}
	newImage, err := ai.NewActionImage(api.cfg, image, action)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create image", http.StatusInternalServerError)
		return
	}

	imagePubID, err := randKey()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to generate ID", http.StatusInternalServerError)
		return
	}

	err = data.SaveImage(context.Background(), imagePubID, newImage)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to save image", http.StatusInternalServerError)
		return
	}

	err = data.AddImageToLastAction(player.ID, imagePubID)
	if err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to add image to last action",
			http.StatusInternalServerError,
		)
		return
	}

	res := ImageActionResponse{
		ImageID: imagePubID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
