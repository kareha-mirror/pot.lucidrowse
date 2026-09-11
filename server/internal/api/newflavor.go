package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type NewFlavorRequest struct {
	Input string `json:"input"`
}

func (api *API) newFlavor(w http.ResponseWriter, r *http.Request) {
	user, err := auth(w, r)
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

	var req NewFlavorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusBadRequest,
			"変な要求です。",
		)
		return
	}

	player, err := data.LoadPlayer(user.ID)
	if err != nil {
		playerPubID, err := randKey()
		if err != nil {
			log.Println(err)
			writeError(
				w,
				http.StatusInternalServerError,
				"約束を決められません。",
			)
			return
		}

		if err = data.CreatePlayer(user.ID, playerPubID); err != nil {
			log.Println(err)
			writeError(
				w,
				http.StatusInternalServerError,
				"あなたの分身を作れません。",
			)
			return
		}

		player, err = data.LoadPlayer(user.ID)
		if err != nil {
			log.Println(err)
			writeError(
				w,
				http.StatusInternalServerError,
				"あなたは分身がいません。",
			)
			return
		}
	}

	flavor, err := ai.NewFlavor(api.cfg, req.Input)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたが何者なのか決められません。",
		)
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

	if err = data.AddFlavor(player.ID, f); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたが何者なのか覚えられません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(flavor)
}
