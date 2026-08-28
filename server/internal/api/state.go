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

func handleRegionState(w http.ResponseWriter, r *http.Request) {
	regionCode := r.PathValue("id")

	areas, err := data.AreaList()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
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
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		b.WriteString(state)

		// TODO
		b.WriteString("\n\n")
	}

	res := RegionStateResponse{
		State: b.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
