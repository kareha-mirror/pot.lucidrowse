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

func (api *API) updateFlavor(w http.ResponseWriter, r *http.Request) {
	user, player, err := authPlayer(w, r)
	if err != nil {
		return
	}

	if user.AICalls >= api.cfg.Game.MaxAICalls {
		writeError(
			w,
			http.StatusTooManyRequests,
			"夢の果実がありません。",
		)
		return
	}
	if err = data.IncrementAICalls(user.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"夢の果実をかじれません。",
		)
		return
	}

	var req UpdateFlavorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusBadRequest,
			"変な要求です。",
		)
		return
	}

	if player.Points < api.cfg.Game.PointsToUpdate {
		log.Println("not enough points")
		writeError(
			w,
			http.StatusBadRequest,
			"書いた回数が足りません。",
		)
		return
	}

	f, err := data.LoadCurrentFlavor(player.ID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusNotFound,
			"あなたはまだ何者なのか決まってません。",
		)
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
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたが何者なのかを改めることができません。",
		)
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
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたが何者なのか覚えられません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedFlavor)
}
