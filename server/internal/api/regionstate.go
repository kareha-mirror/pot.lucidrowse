package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type RegionStateResponse struct {
	State string `json:"state"`
}

func (api *API) regionState(w http.ResponseWriter, r *http.Request) {
	regionCode := r.PathValue("code")

	areas, err := data.AreaList()
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"地域を眺められません。",
		)
		return
	}

	b := strings.Builder{}
	for _, area := range areas {
		if area.RegionCode != regionCode {
			continue
		}

		b.WriteString("== " + area.Name + " ==\n\n")

		state, err := data.AreaState(area.RegionCode + "-" + area.AreaCode)
		if err != nil {
			log.Println(err)
			writeError(
				w,
				http.StatusNotFound,
				"地域の状態が分かりません。",
			)
			return
		}
		b.WriteString(state)

		b.WriteString("\n\n")
	}

	res := RegionStateResponse{
		State: b.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
