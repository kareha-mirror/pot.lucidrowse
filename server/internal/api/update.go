package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

type UpdateRequest struct {
	Id    string `json:"id"`
	Input string `json:"input"`
}

type UpdateResponse struct {
	Id     string `json:"id"`
	Flavor Flavor `json:"flavor"`
}

func update(cfg *config.Config, flavor string, input string) (string, error) {
	return ai.Update(cfg, flavor, input)
}

func handleUpdate(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	var req UpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	flavor, ok := flavors[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	rawUpdatedFlavorStr, err := update(cfg, flavor, req.Input)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var updatedFlavor Flavor

	updatedFlavorStr := strings.ReplaceAll(rawUpdatedFlavorStr, "```json", "")
	updatedFlavorStr = strings.ReplaceAll(updatedFlavorStr, "```", "")
	if err := json.Unmarshal([]byte(updatedFlavorStr), &updatedFlavor); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := CreateResponse{
		Id:     req.Id,
		Flavor: updatedFlavor,
	}

	flavors[req.Id] = updatedFlavorStr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func wrapUpdateHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleUpdate(cfg, w, r)
	}
}
