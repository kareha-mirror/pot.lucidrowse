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

	mux.HandleFunc("GET /api/hello", api.hello)
	mux.HandleFunc("POST /api/next-day", api.nextDay)

	mux.HandleFunc("GET /api/state", api.loadState)
	mux.HandleFunc("GET /api/users", api.loadUser)
	mux.HandleFunc("GET /api/players", api.loadPlayer)

	mux.HandleFunc("POST /api/users/ensure-session", api.ensureSession)

	mux.HandleFunc("POST /api/players/flavor", api.newFlavor)
	mux.HandleFunc("POST /api/players/flavor/image", api.imageFlavor)
	mux.HandleFunc("POST /api/players/flavor/commit", api.commitFlavor)
	mux.HandleFunc("POST /api/players/flavor/update", api.updateFlavor)

	mux.HandleFunc("POST /api/players/actions", api.newAction)
	mux.HandleFunc("POST /api/players/actions/image", api.imageAction)
	mux.HandleFunc("POST /api/players/actions/commit", api.commitAction)

	mux.HandleFunc("GET /api/images/{id}", api.image)

	mux.HandleFunc("GET /api/regions/{code}/players", api.listPlayers)
	mux.HandleFunc("GET /api/players/{id}/actions", api.listActions)
	mux.HandleFunc("GET /api/regions/{code}/state", api.regionState)

	mux.HandleFunc("POST /api/players/release", api.releasePlayer)
	mux.HandleFunc("POST /api/players/{id}/override", api.overridePlayer)

	mux.HandleFunc("POST /api/key/login", api.login)
	mux.HandleFunc("POST /api/key/create", api.createKey)
	mux.HandleFunc("POST /api/key/change", api.changePassword)
	mux.HandleFunc("POST /api/key/logout", api.logout)

	log.Println("Lucidrowse server: " + cfg.App.Addr)
	return http.ListenAndServe(cfg.App.Addr, mux)
}
