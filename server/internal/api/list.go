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

func (api *API) listPlayers(w http.ResponseWriter, r *http.Request) {
	regionCode := r.PathValue("code")

	players, err := data.PlayerList(regionCode)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"住人たちを眺められません。",
		)
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

func (api *API) listActions(w http.ResponseWriter, r *http.Request) {
	playerPubID := r.PathValue("id")

	actions, err := data.ActionList(playerPubID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"日記を眺められません。",
		)
		return
	}

	res := ListActionsResponse{
		Actions: actions,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
