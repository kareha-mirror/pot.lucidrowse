package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type UpdateFlavorRequest struct {
	Input string `json:"input"`
}

func (api *API) handleUpdateFlavor(w http.ResponseWriter, r *http.Request) {
	user, player, err := authPlayer(w, r)
	if err != nil {
		return
	}

	if user.AICalls >= api.cfg.Game.MaxAICalls {
		http.Error(w, "too many requests", http.StatusTooManyRequests)
		return
	}
	if err = data.IncrementAICalls(user.ID); err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to increment AI calls",
			http.StatusInternalServerError,
		)
		return
	}

	var req UpdateFlavorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if player.Points < api.cfg.Game.PointsToUpdate {
		log.Println("not enough points")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if err = data.ResetPlayerPoints(player.ID); err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to reset player points",
			http.StatusInternalServerError,
		)
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

	updatedFlavor, err := ai.UpdateFlavor(api.cfg, flavor, req.Input)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to update flavor", http.StatusInternalServerError)
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

	if err = data.AddFlavor(player.ID, updatedF); err != nil {
		log.Println(err)
		http.Error(w, "failed to add flavor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedFlavor)
}
