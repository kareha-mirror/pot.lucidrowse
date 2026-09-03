package api

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ListPlayersResponse struct {
	Players []data.PlayerItem `json:"players"`
}

func handleListPlayers(w http.ResponseWriter, r *http.Request) {
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

func handleListActions(w http.ResponseWriter, r *http.Request) {
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

func handleListMyActions(w http.ResponseWriter, r *http.Request) {
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

	playerPubID, err := data.PlayerPubID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "player not found", http.StatusNotFound)
		return
	}

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
