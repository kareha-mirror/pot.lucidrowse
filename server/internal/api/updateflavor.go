package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type UpdateFlavorRequest struct {
	PlayerId string `json:"player-id"`
	Input    string `json:"input"`
}

func updateFlavor(
	cfg *config.Config, flavor string, input string,
) (string, error) {
	return ai.Update(cfg, flavor, input)
}

func handleUpdateFlavor(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	var req UpdateFlavorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	id, err := data.PlayerId(req.PlayerId)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	f, err := data.LoadLastFlavor(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
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

	rawUpdatedFlavorStr, err := updateFlavor(cfg, string(flavorStr), req.Input)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	updatedFlavorStr := sanitizeFlavorString(rawUpdatedFlavorStr)
	var updatedFlavor Flavor
	if err := json.Unmarshal([]byte(updatedFlavorStr), &updatedFlavor); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	updatedF := data.Flavor{
		Input:       req.Input,
		Name:        updatedFlavor.Name,
		Race:        updatedFlavor.Race,
		Job:         updatedFlavor.Job,
		Description: updatedFlavor.Description,
		AreaCode:    updatedFlavor.AreaCode,
		AreaName:    updatedFlavor.AreaName,
	}

	if err = data.AddFlavor(id, updatedF); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedFlavor)
}

func updateFlavorHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleUpdateFlavor(cfg, w, r)
	}
}
