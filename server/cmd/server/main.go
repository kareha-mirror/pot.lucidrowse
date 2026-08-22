package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/filter"
)

type HealthResponse struct {
	Awake bool `json:"awake"`
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	res := HealthResponse{
		Awake: true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

type InterpretRequest struct {
	Text string `json:"text"`
}

type Flavor struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Job  string `json:"job"`
	Text string `json:"text"`
}

type InterpretResponse struct {
	Id string `json:"id"`
	Flavor Flavor `json:"flavor"`
}

func interpretCharacter(cfg *config.Config, text string) (string, error) {
	return filter.Apply(cfg, text)
}

var texts  = map[string]string{}

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

	var flavor Flavor

	sanitized := strings.ReplaceAll(result, "```json", "")
	sanitized = strings.ReplaceAll(sanitized, "```", "")
	if err := json.Unmarshal([]byte(sanitized), &flavor); err != nil {
		log.Println(result)
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	id := "test"

	res := InterpretResponse{
		Id: id,
		Flavor: flavor,
	}

	texts[id] = sanitized

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

type ImageRequest struct {
	Id string `json:"id"`
}

func generateImage(cfg *config.Config, text string) ([]byte, error) {
	return filter.ApplyImage(cfg, text)
}

func handleImage(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	var req ImageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	text, ok := texts[req.Id]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	img, err := generateImage(cfg, text)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/webp")
	w.Write(img)
}

func createHandleCharacterInterpret(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleCharacterInterpret(cfg, w, r)
	}
}

func createHandleImage(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleImage(cfg, w, r)
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
	mux.HandleFunc("POST /api/character/interpret", createHandleCharacterInterpret(cfg))
	mux.HandleFunc("POST /api/character/image", createHandleImage(cfg))

	log.Println("Lucidrowse server: " + cfg.App.Addr)
	log.Fatal(http.ListenAndServe(cfg.App.Addr, mux))
}
