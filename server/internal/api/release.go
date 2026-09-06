package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ReleasePlayerResponse struct{}

func (api *API) handleReleasePlayer(w http.ResponseWriter, r *http.Request) {
	user, err := auth(w, r)
	if err != nil {
		return
	}

	err = data.ReleasePlayer(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to release player",
			http.StatusInternalServerError,
		)
		return
	}

	res := ReleasePlayerResponse{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
