package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type ImageActionResponse struct {
	ImageID string `json:"image-id"`
}

func (api *API) imageAction(w http.ResponseWriter, r *http.Request) {
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

	a, err := data.LoadLastAction(player.ID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusNotFound,
			"あなたはまだ何をしたか書いてません。",
		)
		return
	}

	f, err := data.LoadCurrentFlavor(player.ID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusNotFound,
			"あなたはまだ何者なのか決まってません。",
		)
		return
	}

	if f.ImagePubID == nil {
		log.Println("image not found")
		writeError(
			w,
			http.StatusNotFound,
			"あなたは姿がありません。",
		)
		return
	}
	image, err := data.LoadImage(context.Background(), *f.ImagePubID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"あなたは姿が見えません。",
		)
		return
	}

	action := ai.Action{Description: a.Description}
	newImage, err := ai.NewActionImage(api.cfg, image, action)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"情景を描けません。",
		)
		return
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
			"情景を覚えられません。",
		)
		return
	}

	err = data.AddImageToLastAction(player.ID, imagePubID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"情景を日記に描けません。",
		)
		return
	}

	res := ImageActionResponse{
		ImageID: imagePubID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
