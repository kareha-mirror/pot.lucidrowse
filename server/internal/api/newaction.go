package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type NewActionRequest struct {
	Input string `json:"input"`
}

func (api *API) newAction(w http.ResponseWriter, r *http.Request) {
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

	var req NewActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusBadRequest,
			"変な要求です。",
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

	action, err := ai.NewAction(api.cfg, flavor, req.Input)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"何を書くか決められません。",
		)
		return
	}

	if action.Error == "" {
		a := data.Action{
			Input:       req.Input,
			Description: action.Description,
		}

		if err = data.AddAction(player.ID, a); err != nil {
			log.Println(err)
			writeError(
				w,
				http.StatusInternalServerError,
				"日記に書けません。",
			)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(action)
}
