package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ListPlayersResponse struct {
	Players []data.PlayerItem `json:"players"`
}

func handleListPlayers(w http.ResponseWriter, r *http.Request) {
	regionCode := r.PathValue("id")

	players, err := data.PlayerList(regionCode)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := ListPlayersResponse{
		Players: players,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

type ListActionsResponse struct {
	Actions []data.ActionItem `json:"actions"`
}

func handleListActions(w http.ResponseWriter, r *http.Request) {
	playerPubId := r.PathValue("id")

	actions, err := data.ActionList(playerPubId)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := ListActionsResponse{
		Actions: actions,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
