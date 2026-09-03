package api

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type NewActionRequest struct {
	Input string `json:"input"`
}

type NewActionResponse struct {
	Description string `json:"description"`
}

func handleNewAction(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
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

	var req NewActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	player, err := data.LoadPlayer(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "player not found", http.StatusNotFound)
		return
	}

	f, err := data.LoadCurrentFlavor(player.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "current flavor not found", http.StatusNotFound)
		return
	}

	flavor := ai.Flavor{
		Name:        f.Name,
		Race:        f.Race,
		Job:         f.Job,
		Description: f.Description,
		AreaCode:    f.AreaCode,
		AreaName:    f.AreaName,
	}

	action, err := ai.NewAction(cfg, flavor, req.Input)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create action", http.StatusInternalServerError)
		return
	}

	if action.Error == "" {
		a := data.Action{
			Input:       req.Input,
			Description: action.Description,
		}

		if err = data.AddAction(player.ID, a); err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to add action",
				http.StatusInternalServerError,
			)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(action)
}

func newActionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleNewAction(cfg, w, r)
	}
}
