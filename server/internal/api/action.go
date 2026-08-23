package api

import (
	"encoding/json"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

type ActionRequest struct {
	Id    string `json:"id"`
	Input string `json:"input"`
}

type ActionResponse struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func action(cfg *config.Config, flavor string, input string) (string, error) {
	return ai.Action(cfg, flavor, input)
}

var actions = map[string]string{}

func handleAction(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	var req ActionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	flavor, ok := flavors[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	text, err := action(cfg, flavor, req.Input)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := ActionResponse{
		Id:   req.Id,
		Text: text,
	}

	actions[req.Id] = text

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func wrapActionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleAction(cfg, w, r)
	}
}
