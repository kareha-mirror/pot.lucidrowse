package api

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ImageFlavorResponse struct {
	ImageID string `json:"image-id"`
}

func handleImageFlavor(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	keyHash := sha256.Sum256([]byte(cookie.Value))
	user, err := data.LoadUser(keyHash[:])
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if user.AICalls >= cfg.Game.AICalls {
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
	player, err := data.LoadPlayer(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "player not found", http.StatusNotFound)
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
		newImage, err = ai.NewFlavorImage(cfg, flavor)
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
		newImage, err = ai.UpdateFlavorImage(cfg, image, flavor)
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

	imagePubID, err := newPubID()
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

func imageFlavorHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleImageFlavor(cfg, w, r)
	}
}
