package api

import (
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func (api *API) image(w http.ResponseWriter, r *http.Request) {
	imagePubID := r.PathValue("id")

	image, err := data.LoadImage(r.Context(), imagePubID)
	if err != nil {
		log.Println(err)
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", image.ContentType)
	w.Header().Set("Cache-Control", "public, max-age=1209600")
	w.Write(image.Content)
}
