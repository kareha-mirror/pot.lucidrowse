package api

import (
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func handleImage(w http.ResponseWriter, r *http.Request) {
	imagePubId := r.PathValue("id")

	image, err := data.LoadImage(r.Context(), imagePubId)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", image.ContentType)
	w.Write(image.Content)
}
