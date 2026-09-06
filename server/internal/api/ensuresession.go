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

func newSession(cfg *config.Config, w http.ResponseWriter) ([]byte, error) {
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

func (api *API) handleEnsureSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			keyHash, err := newSession(api.cfg, w)
			if err != nil {
				log.Println(err)
				http.Error(
					w,
					"failed to generate session key",
					http.StatusInternalServerError,
				)
				return
			}

			userID, err := data.CreateUser()
			if err != nil {
				log.Println(err)
				http.Error(
					w,
					"failed to create user",
					http.StatusInternalServerError,
				)
				return
			}
			if err = data.AddSession(userID, keyHash); err != nil {
				log.Println(err)
				http.Error(
					w,
					"failed to create session",
					http.StatusInternalServerError,
				)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct{}{})
			return
		}

		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	sessionKey := cookie.Value
	keyHash := sha256.Sum256([]byte(sessionKey))
	_, err = data.LoadUser(keyHash[:])
	if err != nil {
		keyHash, err := newSession(api.cfg, w)
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to generate session key",
				http.StatusInternalServerError,
			)
			return
		}

		userID, err := data.CreateUser()
		if err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to create user",
				http.StatusInternalServerError,
			)
			return
		}
		if err = data.AddSession(userID, keyHash); err != nil {
			log.Println(err)
			http.Error(
				w,
				"failed to create session",
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct{}{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}
