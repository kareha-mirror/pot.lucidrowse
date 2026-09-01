package api

import (
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func Run(cfg *config.Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", handleHello)
	mux.HandleFunc("GET /api/date", handleDate)
	mux.HandleFunc("POST /api/next-day", nextDayHandler(cfg))

	mux.HandleFunc("POST /api/players", handleNewPlayer)

	mux.HandleFunc("POST /api/players/{id}/flavor", newFlavorHandler(cfg))
	mux.HandleFunc(
		"POST /api/players/{id}/flavor/image", imageFlavorHandler(cfg),
	)
	mux.HandleFunc(
		"POST /api/players/{id}/flavor/update", updateFlavorHandler(cfg),
	)
	mux.HandleFunc("POST /api/players/{id}/flavor/commit", handleCommitFlavor)

	mux.HandleFunc("POST /api/players/{id}/actions", newActionHandler(cfg))
	mux.HandleFunc(
		"POST /api/players/{id}/actions/image", imageActionHandler(cfg),
	)
	mux.HandleFunc("POST /api/players/{id}/actions/commit", handleCommitAction)

	mux.HandleFunc("GET /api/images/{id}", handleImage)

	mux.HandleFunc("GET /api/regions/{code}/players", handleListPlayers)
	mux.HandleFunc("GET /api/players/{id}/actions", handleListActions)
	mux.HandleFunc("GET /api/regions/{code}/state", handleRegionState)

	log.Println("Lucidrowse server: " + cfg.App.Addr)
	return http.ListenAndServe(cfg.App.Addr, mux)
}
