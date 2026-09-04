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

type ImageActionResponse struct {
	ImageID string `json:"image-id"`
}

func handleImageAction(
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

	if user.AICalls >= cfg.Game.MaxAICalls {
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
	newImage, err := ai.NewActionImage(cfg, image, action)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create image", http.StatusInternalServerError)
		return
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

func imageActionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleImageAction(cfg, w, r)
	}
}
