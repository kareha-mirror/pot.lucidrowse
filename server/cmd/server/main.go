package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "tea.kareha.org/pot/lucidrowse/server/internal/config"
    "tea.kareha.org/pot/lucidrowse/server/internal/filter"
)

type HealthResponse struct {
    Awake bool `json:"awake"`
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    res := HealthResponse {
        Awake: true,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

type InterpretRequest struct {
    Text string `json:"text"`
}

type InterpretResponse struct {
    Text string `json:"text"`
}

func interpretCharacter(cfg *config.Config, text string) (string, error) {
    return filter.Apply(cfg, text)
}

func handleCharacterInterpret(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
    var req InterpretRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    result, err := interpretCharacter(cfg, req.Text)
    if err != nil {
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    res := InterpretResponse {
        Text: result,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func handler(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        handleCharacterInterpret(cfg, w, r)
    }
}

func main() {
	if len(os.Args) < 2 {
		print("Usage: " + os.Args[0] + " server.yaml\n")
		return
	}
	cfg := config.Load(os.Args[1])

    mux := http.NewServeMux()
    mux.HandleFunc("GET /api/health", handleHealth)
    mux.HandleFunc("POST /api/character/interpret", handler(cfg))

    log.Println("Lucidrowse server: " + cfg.App.Addr)
    log.Fatal(http.ListenAndServe(cfg.App.Addr, mux))
}
