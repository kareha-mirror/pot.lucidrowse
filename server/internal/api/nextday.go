package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

var nextDayMu sync.Mutex

type NextDayResponse struct {
	Day   int64  `json:"day"`
	Error string `json:"error"`
}

func handleNextDay(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	if !nextDayMu.TryLock() {
		res := NextDayResponse{
			Day:   0,
			Error: "excluded",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}
	defer nextDayMu.Unlock()

	areas, err := data.AreaList()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to list areas", http.StatusInternalServerError)
		return
	}

	areaStates := map[string]string{}
	updatedRegions := map[string]struct{}{}
	for _, area := range areas {
		areaCode := area.RegionCode + "-" + area.AreaCode
		eventList, err := data.EventList(areaCode)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to list actions",
				http.StatusInternalServerError,
			)
			return
		}
		if len(eventList) < 1 {
			continue
		}
		state, err := ai.UpdateAreaState(cfg, areaCode)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to update area state",
				http.StatusInternalServerError,
			)
			return
		}
		areaStates[areaCode] = state
		regionCode, err := data.RegionFromArea(areaCode)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to infer region code from area code",
				http.StatusInternalServerError,
			)
			return
		}
		updatedRegions[regionCode] = struct{}{}
	}

	regionStates := map[string]string{}
	for regionCode, _ := range data.RegionNames {
		_, ok := updatedRegions[regionCode]
		if !ok {
			continue
		}

		state, err := ai.UpdateRegionState(cfg, regionCode)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to update region state",
				http.StatusInternalServerError,
			)
			return
		}
		regionStates[regionCode] = state
	}

	worldState, err := ai.UpdateWorldState(cfg)
	if err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to update world state",
			http.StatusInternalServerError,
		)
		return
	}

	day, err := data.NextDay()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to go next day", http.StatusInternalServerError)
		return
	}

	for _, area := range areas {
		areaCode := area.RegionCode + "-" + area.AreaCode
		_, ok := areaStates[areaCode]
		if !ok {
			continue
		}
		err := data.AddAreaState(areaCode, areaStates[areaCode])
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to add area state",
				http.StatusInternalServerError,
			)
			return
		}
	}

	for regionCode, _ := range data.RegionNames {
		regionState, ok := regionStates[regionCode]
		if !ok {
			continue
		}
		err := data.AddRegionState(regionCode, regionState)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to add region state",
				http.StatusInternalServerError,
			)
			return
		}
	}

	err = data.AddWorldState(worldState)
	if err != nil {
		log.Println(err)
		http.Error(
			w,
			"failed to add world state",
			http.StatusInternalServerError,
		)
		return
	}

	res := NextDayResponse{
		Day:   day,
		Error: "",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func nextDayHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleNextDay(cfg, w, r)
	}
}
