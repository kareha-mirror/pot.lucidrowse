package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type CommitFlavorRequest struct {
	PlayerId string `json:"player-id"`
}

type CommitFlavorResponse struct{}

func handleCommitFlavor(w http.ResponseWriter, r *http.Request) {
	var req CommitFlavorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	playerId, err := data.PlayerId(req.PlayerId)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err = data.CommitLastFlavor(playerId); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err = data.ActivatePlayer(playerId); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := CommitFlavorResponse{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
