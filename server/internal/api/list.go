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

func (api *API) handleListPlayers(w http.ResponseWriter, r *http.Request) {
	regionCode := r.PathValue("code")

	players, err := data.PlayerList(regionCode)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to list players", http.StatusInternalServerError)
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

func (api *API) handleListActions(w http.ResponseWriter, r *http.Request) {
	playerPubID := r.PathValue("id")

	actions, err := data.ActionList(playerPubID)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to list actions", http.StatusInternalServerError)
		return
	}

	res := ListActionsResponse{
		Actions: actions,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
