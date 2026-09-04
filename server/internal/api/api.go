package api

import (
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

func Run(cfg *config.Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", handleHello)
	mux.HandleFunc("GET /api/state", stateHandler(cfg))
	mux.HandleFunc("POST /api/next-day", nextDayHandler(cfg))

	mux.HandleFunc("GET /api/users", loadUserHandler(cfg))
	mux.HandleFunc("POST /api/users/ensure-session", ensureSessionHandler(cfg))
	mux.HandleFunc("GET /api/players", handleLoadPlayer)

	mux.HandleFunc("POST /api/players/flavor", newFlavorHandler(cfg))
	mux.HandleFunc(
		"POST /api/players/flavor/image", imageFlavorHandler(cfg),
	)
	mux.HandleFunc(
		"POST /api/players/flavor/update", updateFlavorHandler(cfg),
	)
	mux.HandleFunc("POST /api/players/flavor/commit", handleCommitFlavor)

	mux.HandleFunc("POST /api/players/actions", newActionHandler(cfg))
	mux.HandleFunc(
		"POST /api/players/actions/image", imageActionHandler(cfg),
	)
	mux.HandleFunc("POST /api/players/actions/commit", handleCommitAction)

	mux.HandleFunc("GET /api/images/{id}", handleImage)

	mux.HandleFunc("GET /api/regions/{code}/players", handleListPlayers)
	mux.HandleFunc("GET /api/players/{id}/actions", handleListActions)
	mux.HandleFunc("GET /api/regions/{code}/state", handleRegionState)

	mux.HandleFunc("POST /api/players/release", handleReleasePlayer)
	mux.HandleFunc("POST /api/players/{id}/override", handleOverridePlayer)

	log.Println("Lucidrowse server: " + cfg.App.Addr)
	return http.ListenAndServe(cfg.App.Addr, mux)
}
