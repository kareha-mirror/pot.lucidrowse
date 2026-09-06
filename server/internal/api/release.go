package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func (api *API) releasePlayer(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}
