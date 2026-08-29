package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type CommitActionResponse struct{}

func handleCommitAction(w http.ResponseWriter, r *http.Request) {
	playerPubId := r.PathValue("id")

	playerId, err := data.PlayerId(playerPubId)
	if err != nil {
		log.Println(err)
		http.Error(w, "player not found", http.StatusNotFound)
		return
	}

	if err = data.CommitLastAction(playerId); err != nil {
		log.Println(err)
		http.Error(w, "last action not found", http.StatusNotFound)
		return
	}

	if err = data.ActivatePlayer(playerId); err != nil {
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
