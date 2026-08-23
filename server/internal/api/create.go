package api

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"

	"tea.kareha.org/pot/lucidrowse/server/internal/ai"
	"tea.kareha.org/pot/lucidrowse/server/internal/config"
)

type CreateRequest struct {
	Input string `json:"input"`
}

type Flavor struct {
	Name  string `json:"name"`
	Race  string `json:"race"`
	Job   string `json:"job"`
	Text  string `json:"text"`
	Error string `json:"error"`
}

type CreateResponse struct {
	Id     string `json:"id"`
	Flavor Flavor `json:"flavor"`
}

func create(cfg *config.Config, input string) (string, error) {
	return ai.Create(cfg, input)
}

func random256() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

var flavors = map[string]string{}

func handleCreate(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	var req CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	rawFlavorStr, err := create(cfg, req.Input)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	flavorStr := strings.ReplaceAll(rawFlavorStr, "```json", "")
	flavorStr = strings.ReplaceAll(flavorStr, "```", "")

	var flavor Flavor
	if err := json.Unmarshal([]byte(flavorStr), &flavor); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	id, err := random256()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := CreateResponse{
		Id:     id,
		Flavor: flavor,
	}

	flavors[id] = flavorStr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func wrapCreateHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleCreate(cfg, w, r)
	}
}
