package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func (api *API) handleCommitFlavor(w http.ResponseWriter, r *http.Request) {
	_, player, err := authPlayer(w, r)
	if err != nil {
		return
	}

	if err = data.CommitLastFlavor(player.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusNotFound,
			"あなたはまだ分身を作ってません。",
		)
		return
	}

	if err = data.ActivatePlayer(player.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたはうまく活動できません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}

func (api *API) handleCommitAction(w http.ResponseWriter, r *http.Request) {
	_, player, err := authPlayer(w, r)
	if err != nil {
		return
	}

	if err = data.CommitLastAction(player.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusNotFound,
			"あなたはまだ何をしたか書いてません。",
		)
		return
	}

	if err = data.ActivatePlayer(player.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたはうまく活動できません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}
