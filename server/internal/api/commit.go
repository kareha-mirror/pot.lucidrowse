package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func (api *API) commitFlavor(w http.ResponseWriter, r *http.Request) {
	_, player, err := authPlayer(w, r)
	if err != nil {
		return
	}

	if err = data.CommitLastFlavor(player.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusNotFound,
			"あなたはまだ何者なのか決まってません。",
		)
		return
	}

	if err = data.ResetPlayerPoints(player.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"書いた回数を戻せません。",
		)
		return
	}

	if err = data.ActivatePlayer(player.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたはうまく活動を始められませんでした。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}

func (api *API) commitAction(w http.ResponseWriter, r *http.Request) {
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
			"あなたはうまく活動を始められませんでした。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}
