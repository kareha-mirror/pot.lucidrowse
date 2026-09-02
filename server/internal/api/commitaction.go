package api

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type CommitActionResponse struct{}

func handleCommitAction(w http.ResponseWriter, r *http.Request) {
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

	playerID, err := data.PlayerID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "player not found", http.StatusNotFound)
		return
	}

	if err = data.CommitLastAction(playerID); err != nil {
		log.Println(err)
		http.Error(w, "last action not found", http.StatusNotFound)
		return
	}

	if err = data.ActivatePlayer(playerID); err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to activate player",
			http.StatusInternalServerError,
		)
		return
	}

	res := CommitActionResponse{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
