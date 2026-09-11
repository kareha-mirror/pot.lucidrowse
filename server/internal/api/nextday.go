package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

var nextDayMu sync.Mutex
var nextDayApiMu sync.Mutex

func NextDay(cfg *config.Config) error {
	if !nextDayMu.TryLock() {
		return fmt.Errorf("failed to lock")
	}
	defer nextDayMu.Unlock()

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
	for regionCode := range data.RegionNames {
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

	for regionCode := range data.RegionNames {
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

func (api *API) nextDay(w http.ResponseWriter, r *http.Request) {
	if api.cfg.App.Mode != "devel" {
		log.Println("next day not allowed")
		writeError(
			w,
			http.StatusBadRequest,
			"次の日まで眠れません。",
		)
		return
	}

	if !nextDayApiMu.TryLock() {
		writeError(
			w,
			http.StatusInternalServerError,
			"誰かが眠ってます。",
		)
		return
	}
	defer nextDayApiMu.Unlock()

	if err := NextDay(api.cfg); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"次の日になりませんでした。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(struct{}{})
}
