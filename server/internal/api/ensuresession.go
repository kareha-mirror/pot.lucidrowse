package api

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func newSessionKeyHash(
	cfg *config.Config, w http.ResponseWriter,
) ([]byte, error) {
	sessionKey, err := randKey()
	if err != nil {
		return []byte{}, err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionKey,
		Path:     "/",
		HttpOnly: true,
		Secure:   cfg.App.Mode == "release",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24 * 365,
	})

	keyHash := sha256.Sum256([]byte(sessionKey))
	return keyHash[:], nil
}

func newSession(cfg *config.Config, w http.ResponseWriter) {
	keyHash, err := newSessionKeyHash(cfg, w)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"セッションキーを作れません。",
		)
		return
	}

	userID, err := data.CreateUser()
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"ユーザを作れません。",
		)
		return
	}
	if err = data.AddSession(userID, keyHash); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"セッションを作れません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}

func (api *API) ensureSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			newSession(api.cfg, w)
			return
		}

		log.Println(err)
		writeError(
			w,
			http.StatusBadRequest,
			"変な要求です。",
		)
		return
	}

	sessionKey := cookie.Value
	keyHash := sha256.Sum256([]byte(sessionKey))
	_, err = data.LoadUser(keyHash[:])
	if err != nil {
		newSession(api.cfg, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}
