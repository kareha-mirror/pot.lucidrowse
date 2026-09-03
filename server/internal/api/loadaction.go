package api

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type LoadActionResponse struct {
	Action     ai.Action `json:"action"`
	ImagePubID *string   `json:"image-id"`
}

func handleLoadAction(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	keyHash := sha256.Sum256([]byte(cookie.Value))
	userID, err := data.UserID(keyHash[:])
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id, err := data.PlayerID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "player not found", http.StatusNotFound)
		return
	}

	a, err := data.LoadCurrentAction(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "current action not found", http.StatusNotFound)
		return
	}

	action := ai.Action{
		Description: a.Description,
		Error:       "",
	}

	res := LoadActionResponse{
		Action:     action,
		ImagePubID: a.ImagePubID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
