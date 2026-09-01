package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type DateResponse struct {
	Date string `json:"date"`
}

func handleDate(w http.ResponseWriter, r *http.Request) {
	day, err := data.Day()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to identify day", http.StatusInternalServerError)
		return
	}

	date := data.NewDate(day)

	res := DateResponse{
		Date: date.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
