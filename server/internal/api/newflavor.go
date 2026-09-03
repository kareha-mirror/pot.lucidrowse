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

type NewFlavorRequest struct {
	Input string `json:"input"`
}

func handleNewFlavor(
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

	var req NewFlavorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	playerID, err := data.PlayerID(userID)
	if err != nil {
		playerPubID, err := newPubID()
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to generate ID",
				http.StatusInternalServerError,
			)
			return
		}

		if err = data.CreatePlayer(userID, playerPubID); err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to create player",
				http.StatusInternalServerError,
			)
			return
		}

		playerID, err = data.PlayerID(userID)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to identify player ID",
				http.StatusInternalServerError,
			)
			return
		}
	}

	flavor, err := ai.NewFlavor(cfg, req.Input)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create flavor", http.StatusInternalServerError)
		return
	}

	f := data.Flavor{
		Input:       req.Input,
		Name:        flavor.Name,
		Race:        flavor.Race,
		Job:         flavor.Job,
		Description: flavor.Description,
		AreaCode:    flavor.AreaCode,
		AreaName:    flavor.AreaName,
	}

	if err = data.AddFlavor(playerID, f); err != nil {
		log.Println(err)
		http.Error(w, "failed to add flavor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavor)
}

func newFlavorHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleNewFlavor(cfg, w, r)
	}
}
