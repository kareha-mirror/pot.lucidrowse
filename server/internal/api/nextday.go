package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type NextDayResponse struct {
	Day int64 `json:"day"`
}

func handleNextDay(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	areas, err := data.AreaList()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	states := map[string]string{}
	for _, area := range areas {
		areaCode := area.RegionCode + "-" + area.AreaCode
		eventList, err := data.EventList(areaCode)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		if len(eventList) < 1 {
			continue
		}
		state, err := ai.UpdateState(cfg, areaCode)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		states[areaCode] = state
	}

	day, err := data.NextDay()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	for _, area := range areas {
		areaCode := area.RegionCode + "-" + area.AreaCode
		_, ok := states[areaCode]
		if !ok {
			continue
		}
		err := data.AddAreaState(areaCode, states[areaCode])
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	res := NextDayResponse{
		Day: day,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func nextDayHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleNextDay(cfg, w, r)
	}
}
