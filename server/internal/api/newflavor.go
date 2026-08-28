package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type NewFlavorRequest struct {
	PlayerId string `json:"player-id"`
	Input    string `json:"input"`
}

func sanitizeFlavorString(s string) string {
	sanitized := strings.ReplaceAll(s, "```json", "")
	return strings.ReplaceAll(sanitized, "```", "")
}

func handleNewFlavor(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	var req NewFlavorRequest
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

	flavor, err := ai.NewFlavor(cfg, req.Input)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
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

	if err = data.AddFlavor(playerId, f); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
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
