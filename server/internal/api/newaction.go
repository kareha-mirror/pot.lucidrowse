package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type NewActionRequest struct {
	PlayerId string `json:"player-id"`
	Input    string `json:"input"`
}

type NewActionResponse struct {
	Description string `json:"description"`
}

func newAction(
	cfg *config.Config, flavor string, input string,
) (string, error) {
	return ai.Action(cfg, flavor, input)
}

func handleNewAction(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	var req NewActionRequest
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

	f, err := data.LoadLastFlavor(playerId)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	flavor := Flavor{
		Name:        f.Name,
		Race:        f.Race,
		Job:         f.Job,
		Description: f.Description,
		AreaCode:    f.AreaCode,
		AreaName:    f.AreaName,
	}

	flavorStr, err := json.Marshal(flavor)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	description, err := newAction(cfg, string(flavorStr), req.Input)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	a := data.Action{
		Input:       req.Input,
		Description: description,
	}

	if err = data.AddAction(playerId, a); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := NewActionResponse{
		Description: description,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func newActionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleNewAction(cfg, w, r)
	}
}
