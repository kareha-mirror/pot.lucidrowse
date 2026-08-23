package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

type UpdateRequest struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type UpdateResponse struct {
	Id     string `json:"id"`
	Flavor Flavor `json:"flavor"`
}

func update(cfg *config.Config, current string, text string) (string, error) {
	return ai.Update(cfg, current, text)
}

func handleUpdate(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
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

	result, err := update(cfg, current, req.Text)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var flavor Flavor

	sanitized := strings.ReplaceAll(result, "```json", "")
	sanitized = strings.ReplaceAll(sanitized, "```", "")
	if err := json.Unmarshal([]byte(sanitized), &flavor); err != nil {
		log.Println(result)
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := CreateResponse{
		Id:     req.Id,
		Flavor: flavor,
	}

	texts[req.Id] = sanitized

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func wrapUpdateHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleUpdate(cfg, w, r)
	}
}
