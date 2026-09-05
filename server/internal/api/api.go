package api

import (
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

type API struct {
	cfg *config.Config
}

func Run(cfg *config.Config) error {
	mux := http.NewServeMux()

	api := &API{cfg: cfg}

	mux.HandleFunc("GET /api/hello", api.handleHello)
	mux.HandleFunc("GET /api/state", api.handleState)
	mux.HandleFunc("POST /api/next-day", api.handleNextDay)

	mux.HandleFunc("GET /api/users", api.handleLoadUser)
	mux.HandleFunc("POST /api/users/ensure-session", api.handleEnsureSession)
	mux.HandleFunc("GET /api/players", api.handleLoadPlayer)

	mux.HandleFunc("POST /api/players/flavor", api.handleNewFlavor)
	mux.HandleFunc("POST /api/players/flavor/image", api.handleImageFlavor)
	mux.HandleFunc("POST /api/players/flavor/update", api.handleUpdateFlavor)
	mux.HandleFunc("POST /api/players/flavor/commit", api.handleCommitFlavor)

	mux.HandleFunc("POST /api/players/actions", api.handleNewAction)
	mux.HandleFunc("POST /api/players/actions/image", api.handleImageAction)
	mux.HandleFunc("POST /api/players/actions/commit", api.handleCommitAction)

	mux.HandleFunc("GET /api/images/{id}", api.handleImage)

	mux.HandleFunc("GET /api/regions/{code}/players", api.handleListPlayers)
	mux.HandleFunc("GET /api/players/{id}/actions", api.handleListActions)
	mux.HandleFunc("GET /api/regions/{code}/state", api.handleRegionState)

	mux.HandleFunc("POST /api/players/release", api.handleReleasePlayer)
	mux.HandleFunc("POST /api/players/{id}/override", api.handleOverridePlayer)

	log.Println("Lucidrowse server: " + cfg.App.Addr)
	return http.ListenAndServe(cfg.App.Addr, mux)
}
