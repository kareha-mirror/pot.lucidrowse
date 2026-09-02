package api

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

type NewPlayerResponse struct {
	PlayerID string `json:"player-id"`
}

func newPubID() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func handleNewPlayer(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	keyHash := sha256.Sum256([]byte(cookie.Value))
	userID, err := data.UserID(keyHash[:])
	if err != nil {
		log.Println(err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	playerID, err := newPubID()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to generate ID", http.StatusInternalServerError)
		return
	}

	if err = data.CreatePlayer(userID, playerID); err != nil {
		log.Println(err)
		http.Error(w, "failed to create player", http.StatusInternalServerError)
		return
	}

	res := NewPlayerResponse{
		PlayerID: playerID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
