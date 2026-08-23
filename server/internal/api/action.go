package api

import (
	"encoding/json"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

type ActionRequest struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type ActionResponse struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func action(cfg *config.Config, current string, text string) (string, error) {
	return ai.Action(cfg, current, text)
}

var actions = map[string]string{}

func handleAction(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	var req UpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	current, ok := texts[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	result, err := action(cfg, current, req.Text)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := ActionResponse{
		Id:   req.Id,
		Text: result,
	}

	actions[req.Id] = result

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func wrapActionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleAction(cfg, w, r)
	}
}
