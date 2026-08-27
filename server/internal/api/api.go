package api

import (
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func Run(cfg *config.Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", handleHello)
	mux.HandleFunc("GET /api/day", handleDay)
	mux.HandleFunc("GET /api/next-day", handleNextDay)

	mux.HandleFunc("GET /api/new-player", handleNewPlayer)

	mux.HandleFunc("POST /api/new-flavor", newFlavorHandler(cfg))
	mux.HandleFunc("POST /api/image-flavor", imageFlavorHandler(cfg))
	mux.HandleFunc("POST /api/update-flavor", updateFlavorHandler(cfg))
	mux.HandleFunc("POST /api/commit-flavor", handleCommitFlavor)

	mux.HandleFunc("POST /api/new-action", newActionHandler(cfg))
	mux.HandleFunc("POST /api/image-action", imageActionHandler(cfg))
	mux.HandleFunc("POST /api/commit-action", handleCommitAction)

	mux.HandleFunc("GET /api/image/{id}", handleImage)

	mux.HandleFunc("GET /api/list-players/{id}", handleListPlayers)
	mux.HandleFunc("GET /api/list-actions/{id}", handleListActions)

	log.Println("Lucidrowse server: " + cfg.App.Addr)
	return http.ListenAndServe(cfg.App.Addr, mux)
}
