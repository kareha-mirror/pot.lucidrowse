package api

import (
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func Run(cfg *config.Config) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/hello", handleHello)
	mux.HandleFunc("POST /api/player/create", wrapCreateHandler(cfg))
	mux.HandleFunc("POST /api/player/image", wrapImageHandler(cfg))
	mux.HandleFunc("POST /api/player/update", wrapUpdateHandler(cfg))
	mux.HandleFunc("POST /api/player/action", wrapActionHandler(cfg))
	mux.HandleFunc("POST /api/player/action-image", wrapActionImageHandler(cfg))

	log.Println("Lucidrowse server: " + cfg.App.Addr)
	log.Fatal(http.ListenAndServe(cfg.App.Addr, mux))
}
