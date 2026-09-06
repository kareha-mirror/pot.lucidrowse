package api

import (
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type StateResponse struct {
	Mode           string `json:"mode"`
	MaxAICalls     int    `json:"max-ai-calls"`
	PointsToUpdate int64  `json:"points-to-update"`
	Day            int64  `json:"day"`
}

func (api *API) handleState(w http.ResponseWriter, r *http.Request) {
	day, err := data.Day()
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"今日が何日なのか分かりません。",
		)
		return
	}

	res := StateResponse{
		Mode:           api.cfg.App.Mode,
		MaxAICalls:     api.cfg.Game.MaxAICalls,
		PointsToUpdate: api.cfg.Game.PointsToUpdate,
		Day:            day,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
