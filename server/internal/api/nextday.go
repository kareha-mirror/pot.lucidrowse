package api

import (
	"encoding/json"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type NextDayResponse struct {
	Day int64 `json:"day"`
}

func handleNextDay(w http.ResponseWriter, r *http.Request) {
	dayCount, err := data.NextDay()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := NextDayResponse{
		Day: dayCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
