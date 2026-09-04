package api

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type CommitFlavorResponse struct{}

func handleCommitFlavor(w http.ResponseWriter, r *http.Request) {
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

	player, err := data.LoadPlayer(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "player not found", http.StatusNotFound)
		return
	}

	if err = data.CommitLastFlavor(player.ID); err != nil {
		log.Println(err)
		http.Error(w, "last flavor not found", http.StatusNotFound)
		return
	}

	if err = data.ActivatePlayer(player.ID); err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to activate player",
			http.StatusInternalServerError,
		)
		return
	}

	res := CommitFlavorResponse{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
