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
	Error string `json:"error"`
}

func nextDay(cfg *config.Config) error {
	areas, err := data.AreaList()
	if err != nil {
		log.Println(err)
		return err
	}

	areaStates := map[string]string{}
	updatedRegions := map[string]struct{}{}
	for _, area := range areas {
		areaCode := area.RegionCode + "-" + area.AreaCode
		eventList, err := data.EventList(areaCode)
		if err != nil {
			log.Println(err)
			return err
		}
		if len(eventList) < 1 {
			continue
		}
		state, err := ai.UpdateAreaState(cfg, areaCode)
		if err != nil {
			log.Println(err)
			return err
		}
		areaStates[areaCode] = state
		regionCode, err := data.RegionFromArea(areaCode)
		if err != nil {
			log.Println(err)
			return err
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
			return err
		}
		regionStates[regionCode] = state
	}

	var worldState string
	if len(updatedRegions) > 0 {
		worldState, err = ai.UpdateWorldState(cfg)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	_, err = data.NextDay()
	if err != nil {
		log.Println(err)
		return err
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
			return err
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
			return err
		}
	}

	if len(updatedRegions) > 0 {
		err = data.AddWorldState(worldState)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func handleNextDay(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	if !nextDayMu.TryLock() {
		res := NextDayResponse{
			Error: "excluded",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}
	defer nextDayMu.Unlock()

	if err := nextDay(cfg); err != nil {
		log.Println(err)
		http.Error(w, "failed to go next day", http.StatusInternalServerError)
		return
	}

	res := NextDayResponse{
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
