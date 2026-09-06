package api

import (
	"encoding/json"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type LoadUserResponse struct {
	Authorized bool    `json:"authorized"`
	Name       *string `json:"name"`
	AICalls    int     `json:"ai-calls"`
}

func (api *API) handleLoadUser(w http.ResponseWriter, r *http.Request) {
	user, err := auth(w, r)
	if err != nil {
		return
	}

	res := LoadUserResponse{
		Authorized: true,
		Name:       user.Name,
		AICalls:    user.AICalls,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

type LoadPlayerResponse struct {
	PubID            *string    `json:"id"`
	Day              *int64     `json:"day"`
	Flavor           *ai.Flavor `json:"flavor"`
	FlavorImagePubID *string    `json:"flavor-image-id"`
	Action           *ai.Action `json:"action"`
	ActionImagePubID *string    `json:"action-image-id"`
	Points           int64      `json:"points"`
}

func (api *API) handleLoadPlayer(w http.ResponseWriter, r *http.Request) {
	_, player, err := authPlayer(w, r)
	if err != nil {
		return
	}

	var res LoadPlayerResponse
	res.PubID = &player.PubID
	res.Day = &player.Day
	res.Points = player.Points

	f, err := data.LoadCurrentFlavor(player.ID)
	if err == nil {
		flavor := ai.Flavor{
			Name:        f.Name,
			Race:        f.Race,
			Job:         f.Job,
			Description: f.Description,
			AreaCode:    f.AreaCode,
			AreaName:    f.AreaName,
			Error:       "",
		}
		res.Flavor = &flavor
		res.FlavorImagePubID = f.ImagePubID
	}

	a, err := data.LoadCurrentAction(player.ID)
	if err == nil {
		action := ai.Action{
			Description: a.Description,
			Error:       "",
		}
		res.Action = &action
		res.ActionImagePubID = a.ImagePubID
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
