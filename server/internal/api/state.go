package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type StateResponse struct {
	Mode string `json:"mode"`
	Day  int64  `json:"day"`
}

func handleState(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	day, err := data.Day()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to identify day", http.StatusInternalServerError)
		return
	}

	res := StateResponse{
		Mode: cfg.App.Mode,
		Day:  day,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func stateHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleState(cfg, w, r)
	}
}
