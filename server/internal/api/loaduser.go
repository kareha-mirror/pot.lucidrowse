package api

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type LoadUserResponse struct {
	Authorized bool    `json:"authorized"`
	Name       *string `json:"name"`
	AICalls    int     `json:"ai-calls"`
}

func handleLoadUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		res := LoadUserResponse{Authorized: false}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	keyHash := sha256.Sum256([]byte(cookie.Value))
	user, err := data.LoadUser(keyHash[:])
	if err != nil {
		res := LoadUserResponse{Authorized: false}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	res := LoadUserResponse{
		Authorized: true,
		Name:       user.Name,
		AICalls:    user.AICalls,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
