package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

var overridePlayerMu sync.Mutex

func (api *API) overridePlayer(w http.ResponseWriter, r *http.Request) {
	if !overridePlayerMu.TryLock() {
		writeError(
			w,
			http.StatusInternalServerError,
			"別の誰かの順番です。",
		)
		return
	}
	defer overridePlayerMu.Unlock()

	playerPubID := r.PathValue("id")

	user, err := auth(w, r)
	if err != nil {
		return
	}

	err = data.ReleasePlayer(user.ID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"自分を手放せません。",
		)
		return
	}

	err = data.OverridePlayer(user.ID, playerPubID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusBadRequest,
			"入り込めません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}
