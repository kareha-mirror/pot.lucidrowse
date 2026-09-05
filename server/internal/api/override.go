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

func (api *API) handleOverridePlayer(w http.ResponseWriter, r *http.Request) {
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
	user, err := data.LoadUser(keyHash[:])
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err = data.ReleasePlayer(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to release player",
			http.StatusInternalServerError,
		)
		return
	}

	player, err := data.LoadPlayer(user.ID)
	if err == nil && player.Activated {
		log.Println("player found")
		http.Error(w, "player found", http.StatusBadRequest)
		return
	}

	err = data.OverridePlayer(user.ID, playerPubID)
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
