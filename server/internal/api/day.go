package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type DayResponse struct {
	Day int64 `json:"day"`
}

func handleDay(w http.ResponseWriter, r *http.Request) {
	day, err := data.Day()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := DayResponse{
		Day: day,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
