package api

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

var overridePlayerMu sync.Mutex

type OverridePlayerResponse struct {
	Error string `json:"error"`
}

func handleOverridePlayer(w http.ResponseWriter, r *http.Request) {
	if !overridePlayerMu.TryLock() {
		res := OverridePlayerResponse{
			Error: "excluded",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}
	defer overridePlayerMu.Unlock()

	playerPubID := r.PathValue("id")

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

	player, err := data.LoadPlayer(userID)
	if err == nil && player.Activated {
		log.Println("player found")
		http.Error(w, "player found", http.StatusBadRequest)
		return
	}

	err = data.OverridePlayer(userID, playerPubID)
	if err != nil {
		log.Println("override")
		log.Println(err)
		http.Error(w, "failed to override player", http.StatusBadRequest)
		return
	}

	var res OverridePlayerResponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
