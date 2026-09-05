package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type StateResponse struct {
	Mode string `json:"mode"`
	Day  int64  `json:"day"`
}

func (api *API) handleState(w http.ResponseWriter, r *http.Request) {
	day, err := data.Day()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to identify day", http.StatusInternalServerError)
		return
	}

	res := StateResponse{
		Mode: api.cfg.App.Mode,
		Day:  day,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
