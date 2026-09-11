package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ImageFlavorResponse struct {
	ImageID string `json:"image-id"`
}

func (api *API) imageFlavor(w http.ResponseWriter, r *http.Request) {
	user, player, err := authPlayer(w, r)
	if err != nil {
		return
	}

	if user.AICalls >= api.cfg.Game.MaxAICalls {
		writeError(
			w,
			http.StatusBadRequest,
			"夢の果実がありません。",
		)
		return
	}
	if err = data.IncrementAICalls(user.ID); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"夢の果実をかじれません。",
		)
		return
	}

	f, err := data.LoadLastFlavor(player.ID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusNotFound,
			"あなたはまだ何者なのか決まってません。",
		)
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

	var newImage data.Image
	if f.ImagePubID == nil {
		newImage, err = ai.NewFlavorImage(api.cfg, flavor)
		if err != nil {
			log.Println(err)
			writeError(
				w,
				http.StatusInternalServerError,
				"あなたの姿を描けません。",
			)
			return
		}
	} else {
		image, err := data.LoadImage(context.Background(), *f.ImagePubID)
		if err != nil {
			log.Println(err)
			writeError(
				w,
				http.StatusInternalServerError,
				"あなたの姿を思い出せません。",
			)
			return
		}
		newImage, err = ai.UpdateFlavorImage(api.cfg, image, flavor)
		if err != nil {
			log.Println(err)
			writeError(
				w,
				http.StatusInternalServerError,
				"あなたの姿を改められません。",
			)
			return
		}
	}

	imagePubID, err := randKey()
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"約束を決められません。",
		)
		return
	}

	err = data.SaveImage(context.Background(), imagePubID, newImage)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたの姿を覚えられません。",
		)
		return
	}

	err = data.AddImageToLastFlavor(player.ID, imagePubID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたの姿を決められません。",
		)
		return
	}

	res := ImageFlavorResponse{
		ImageID: imagePubID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(res)
}
