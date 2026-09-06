package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ImageFlavorResponse struct {
	ImageID string `json:"image-id"`
}

func (api *API) handleImageFlavor(w http.ResponseWriter, r *http.Request) {
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

	f, err := data.LoadLastFlavor(player.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "last flavor not found", http.StatusNotFound)
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
	if f.ImagePubID == nil {
		newImage, err = ai.NewFlavorImage(api.cfg, flavor)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to create image",
				http.StatusInternalServerError,
			)
			return
		}
	} else {
		image, err := data.LoadImage(context.Background(), *f.ImagePubID)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to load image",
				http.StatusInternalServerError,
			)
			return
		}
		newImage, err = ai.UpdateFlavorImage(api.cfg, image, flavor)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to update image",
				http.StatusInternalServerError,
			)
			return
		}
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

	err = data.AddImageToLastFlavor(player.ID, imagePubID)
	if err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to add image to last flavor",
			http.StatusInternalServerError,
		)
		return
	}

	res := ImageFlavorResponse{
		ImageID: imagePubID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
