package api

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type UpdateFlavorRequest struct {
	Input string `json:"input"`
}

func handleUpdateFlavor(
	cfg *config.Config, w http.ResponseWriter, r *http.Request,
) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	keyHash := sha256.Sum256([]byte(cookie.Value))
	userID, err := data.UserID(keyHash[:])
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req UpdateFlavorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	id, err := data.PlayerID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "player not found", http.StatusNotFound)
		return
	}

	f, err := data.LoadCurrentFlavor(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "current flavor not found", http.StatusNotFound)
		return
	}

	flavor := ai.Flavor{
		Name:        f.Name,
		Race:        f.Race,
		Job:         f.Job,
		Description: f.Description,
		AreaCode:    f.AreaCode,
		AreaName:    f.AreaName,
	}

	updatedFlavor, err := ai.UpdateFlavor(cfg, flavor, req.Input)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to update flavor", http.StatusInternalServerError)
		return
	}

	updatedF := data.Flavor{
		Input:       req.Input,
		Name:        updatedFlavor.Name,
		Race:        updatedFlavor.Race,
		Job:         updatedFlavor.Job,
		Description: updatedFlavor.Description,
		AreaCode:    updatedFlavor.AreaCode,
		AreaName:    updatedFlavor.AreaName,
	}

	if err = data.AddFlavor(id, updatedF); err != nil {
		log.Println(err)
		http.Error(w, "failed to add flavor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedFlavor)
}

func updateFlavorHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleUpdateFlavor(cfg, w, r)
	}
}
