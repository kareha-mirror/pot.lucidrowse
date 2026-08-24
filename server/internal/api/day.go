package api

import (
	"encoding/json"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type DayResponse struct {
	Day int64 `json:"day"`
}

func handleDay(w http.ResponseWriter, r *http.Request) {
	dayCount, err := data.Day()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := DayResponse{
		Day: dayCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
