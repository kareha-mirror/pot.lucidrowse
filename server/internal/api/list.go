package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
	w.Header().Set("Cache-Control", "no-store")
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
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(res)
}

type RegionsResponse struct {
	PlayerCounts map[string]int    `json:"player-counts"`
	Topics       map[string]string `json:"topics"`
}

func (api *API) regions(w http.ResponseWriter, r *http.Request) {
	playerCounts, err := data.PlayerCounts()
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"住人の数が分かりません。",
		)
		return
	}

	states, err := data.RegionStates()
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"地方の状態が分かりません。",
		)
		return
	}

	topics := map[string]string{}
loop:
	for code, state := range states {
		state = strings.ReplaceAll(state, "\r", "")
		parts := strings.Split(state, "\n")
		for _, part := range parts {
			if strings.HasPrefix(part, "- ") {
				topics[code] = part[2:]
				continue loop
			}
		}
	}

	res := RegionsResponse{
		PlayerCounts: playerCounts,
		Topics:       topics,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(res)
}
