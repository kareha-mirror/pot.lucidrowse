package api

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type LoadPlayerResponse struct {
	PubID            *string    `json:"id"`
	Day              *int64     `json:"day"`
	Flavor           *ai.Flavor `json:"flavor"`
	FlavorImagePubID *string    `json:"flavor-image-id"`
	Action           *ai.Action `json:"action"`
	ActionImagePubID *string    `json:"action-image-id"`
}

func handleLoadPlayer(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		res := LoadPlayerResponse{}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	keyHash := sha256.Sum256([]byte(cookie.Value))
	userID, err := data.UserID(keyHash[:])
	if err != nil {
		res := LoadPlayerResponse{}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	player, err := data.LoadPlayer(userID)
	if err != nil {
		res := LoadPlayerResponse{}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	var res LoadPlayerResponse
	res.PubID = &player.PubID
	res.Day = &player.Day

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
